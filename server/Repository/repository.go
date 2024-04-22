package repository

import (
	"fmt"
	"log"
	"time"

	pg "github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"

	model "github.com/khyatidoshi/ShortURLGenerator/server/Model"
)

type URLRepository struct {
	Postgres *pg.DB
	// Redis    *redis.Client
}

func NewURLRepository() *URLRepository {
	pgdb := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "docker",
		Database: "url_generator",
		Addr:     "postgres:5432",
	})
	// Check if the connection was successful
	if pgdb == nil {
		log.Print("failed to connect to the Postgres database")
	}

	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// 	// Add other configuration options
	// })

	// _, err := redisClient.Ping().Result()
	// if err != nil {
	// 	log.Println("Failed to connect to Redis:", err)
	// }

	createSchema(pgdb)

	return &URLRepository{
		Postgres: pgdb,
		// Redis:    redisClient,
	}
}

func createSchema(db *pg.DB) {
	models := []interface{}{(*model.UrlData)(nil)}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			log.Println(err)
		}
	}

	models = []interface{}{(*model.AccessDetails)(nil)}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			log.Println(err)
		}
	}

	_, err := db.Exec(`
		ALTER TABLE eventdetails 
		ADD CONSTRAINT fk_short_url 
		FOREIGN KEY (short_url ) 
		REFERENCES tinyurldata(short_url)
	`)
	if err != nil {
		log.Println(err)
	}
}

func (repo *URLRepository) StoreURL(urlData *model.UrlData) error {
	_, err := repo.Postgres.Model(urlData).Insert()
	if err != nil {
		return fmt.Errorf("failed to store URL: %v", err)
	}

	// Cache the long URL and its corresponding short URL in Redis
	// err = repo.Redis.Set(urlData.ShortURL, urlData.LongURL, 0).Err()
	// if err != nil {
	// 	log.Println("Failed to cache URL:", err)
	// }

	return nil
}

func (repo *URLRepository) FetchURL(shortURL string) (string, error) {
	// longURL, err := repo.Redis.Get(shortURL).Result()
	// if err == nil {
	// 	return longURL, nil
	// }

	urlData := new(model.UrlData)
	err := repo.Postgres.Model(urlData).Where("short_url = ?", shortURL).Select()
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %v", err)
	}

	// err = repo.Redis.Set(shortURL, urlData.LongURL, 0).Err()
	// if err != nil {
	// 	log.Println("Failed to cache URL:", err)
	// }

	return urlData.LongURL, nil
}

func (repo *URLRepository) RecordAccessEvent(shortURL string) error {
	event := &model.AccessDetails{
		ShortURL:   shortURL,
		AccessedAt: time.Now().Unix(),
	}

	_, err := repo.Postgres.Model(event).Insert()
	if err != nil {
		return err
	}

	return nil
}

func (repo *URLRepository) GetAccessCounts(shortURL string) (model.Stats, error) {
	stats := model.Stats{}

	now := time.Now().Unix()
	last24Hours := now - 24*60*60
	pastWeek := now - 7*24*60*60

	query := `
        SELECT 
            COUNT(*) FILTER (WHERE accessed_at >= ?) AS day_count,
            COUNT(*) FILTER (WHERE accessed_at >= ?) AS week_count,
            COUNT(*) AS total_count
        FROM eventdetails
        WHERE short_url = ?
    `

	_, err := repo.Postgres.QueryOne(&stats, query, last24Hours, pastWeek, shortURL)
	if err != nil {
		return stats, err
	}

	return stats, nil
}

func (repo *URLRepository) DeleteExpiredURLs() error {
	var expiredURLs []model.UrlData
	err := repo.Postgres.Model(&expiredURLs).Where("expiry < ?", time.Now().Unix()).Select()
	if err != nil {
		return fmt.Errorf("failed to fetch expired URLs: %v", err)
	}

	// Extract short URLs from the fetched data
	var shortURLs []string
	for _, url := range expiredURLs {
		shortURLs = append(shortURLs, url.ShortURL)
	}

	if len(shortURLs) < 1 {
		return nil
	}

	// for _, shortURL := range shortURLs {
	// 	err := repo.Redis.Del(shortURL).Err()
	// 	if err != nil {
	// 		log.Printf("Failed to delete expired URL %s from Redis cache: %v", shortURL, err)
	// 	}
	// }

	query := "SELECT short_url FROM short_url WHERE short_url IN ?"
	_, err = repo.Postgres.Exec(query, shortURLs)
	if err != nil {
		return err
	}

	query = "DELETE FROM tinyurldata WHERE short_url IN ?"
	_, err = repo.Postgres.Exec(query, shortURLs)
	if err != nil {
		return err
	}

	return nil
}
