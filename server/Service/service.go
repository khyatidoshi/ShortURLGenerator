package service

import (
	"fmt"

	model "github.com/khyatidoshi/ShortURLGenerator/server/Model"
	repo "github.com/khyatidoshi/ShortURLGenerator/server/Repository"
	utils "github.com/khyatidoshi/ShortURLGenerator/server/Utils"
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
		ShortURL:   shortURL,
		LongURL:    req.LongURL,
		ExpiryDate: expiry,
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
		return "", err
	}

	return longURL, nil
}

func (svc *URLService) DeleteURL() error {
	return svc.URLRepo.DeleteExpiredURLs()
}
