package service

import (
	"context"
	"fmt"
	"log"
	"time"

	model "github.com/khyatidoshi/ShortURLGenerator/server/Model"
	repo "github.com/khyatidoshi/ShortURLGenerator/server/Repository"
	utils "github.com/khyatidoshi/ShortURLGenerator/server/Utils"
	"github.com/segmentio/kafka-go"
)

type URLService struct {
	URLRepo *repo.URLRepository
}

func NewURLService() *URLService {
	return &URLService{
		URLRepo: repo.NewURLRepository(),
	}
}

func (svc *URLService) ShortenURL(req model.GenerateShortURLReq, expiry int64) (string, error) {
	shortURL := utils.GenerateShortURL()

	urlData := &model.UrlData{
		ShortURL: shortURL,
		LongURL:  req.LongURL,
	}

	if expiry != 0 {
		urlData.ExpiryDate = expiry
	}

	err := svc.URLRepo.StoreURL(urlData)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://localhost:4000/%s", shortURL), nil
}

func (svc *URLService) GetLongURL(shortURL string) (string, error) {
	longURL, err := svc.URLRepo.FetchURL(shortURL)
	if err != nil {
		return longURL, err
	}
	return longURL, nil
}
func (svc *URLService) RecordURLAccess(shortURL string) {
	err := svc.URLRepo.RecordAccessEvent(shortURL)
	if err != nil {
		fmt.Printf("failed to store event for : %s at %s with error %s", shortURL, time.Now(), err)
	}
}

func (svc *URLService) GetAccessCounts(shortURL string) (model.Stats, error) {
	accessCounts, err := svc.URLRepo.GetAccessCounts(shortURL)
	if err != nil {
		return accessCounts, err
	}

	return accessCounts, nil
}

func (svc *URLService) DeleteURL() error {
	return svc.URLRepo.DeleteExpiredURLs()
}

// Kafka consumer setup
func (svc *URLService) ConsumeMessage() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"kafka:9092"},
		Topic:          "url-redirects",
		GroupID:        "url-redirect-group",
		MinBytes:       10e3,        // 10KB
		MaxBytes:       10e6,        // 10MB
		CommitInterval: time.Second, // frequency of offset commit
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error while receiving message: %s\n", err)
			continue
		}
		shortURL := string(m.Value)
		log.Printf("Processing URL: %s\n", shortURL)
		err = svc.URLRepo.RecordAccessEvent(shortURL)
		if err != nil {
			log.Printf("failed to recover event for url : %s with err: %s, at %d", shortURL, err, time.Now().Unix())
		}
	}
}
