package repository

import (
	"fmt"

	pg "github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"

	model "github.com/khyatidoshi/ShortURLGenerator/server/Model"
)

type URLRepository struct {
	Postgres *pg.DB
}

func NewURLRepository() *URLRepository {
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// database := os.Getenv("DB_NAME")
	// host := os.Getenv("DB_HOST")

	pgdb := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
		Addr:     "localhost:5432",
	})

	// Check if the connection was successful
	if pgdb == nil {
		fmt.Print("failed to connect to the Postgres database")
	}

	createSchema(pgdb)

	return &URLRepository{Postgres: pgdb}
}

func createSchema(db *pg.DB) {
	models := []interface{}{(*model.UrlData)(nil)}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (repo *URLRepository) StoreURL(urlData *model.UrlData) error {
	_, err := repo.Postgres.Model(urlData).Insert()

	if err != nil {
		return fmt.Errorf("failed to store URL: %v", err)
	}
	return nil
}

func (repo *URLRepository) FetchURL(shortURL string) (string, error) {
	urlData := new(model.UrlData)

	err := repo.Postgres.Model(urlData).Where("short_url = ?", shortURL).Select()
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %v", err)
	}

	return urlData.LongURL, nil
}
