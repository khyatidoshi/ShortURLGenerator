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

func (svc *URLService) ShortenURL(req model.GenerateShortURLReq) (string, error) {
	shortURL := utils.GenerateShortURL()

	urlData := &model.UrlData{
		ShortURL:   shortURL,
		LongURL:    req.LongURL,
		ExpiryDate: req.ExpiryDate,
	}

	err := svc.URLRepo.StoreURL(urlData)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://localhost:8080/", shortURL), nil
}

func (svc *URLService) GetLongURL(shortURL string) (string, error) {
	longURL, err := svc.URLRepo.FetchURL(shortURL)
	if err != nil {
		return "", err
	}

	return longURL, nil
}

func (svc *URLService) DeleteURL(shortURL string) error {
	// err := svc.URLRepo.DeleteURL(shortURL)
	// if err != nil {
	// 	return err
	// }
	return nil
}
