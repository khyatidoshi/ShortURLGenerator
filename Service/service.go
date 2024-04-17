package service

import (
	repo "github.com/khyatidoshi/ShourtURLGenerator/Repository"
	utils "github.com/khyatidoshi/ShourtURLGenerator/Utils"
)

type URLService struct {
	URLRepo *repo.URLRepository
}

func NewURLService() *URLService {
	return &URLService{
		URLRepo: repo.NewURLRepository(),
	}
}

func (svc *URLService) ShortenURL(longURL string) (string, error) {
	shortURL := utils.GenerateShortURL()
	// err := svc.URLRepo.StoreURL(shortURL, longURL)
	// if err != nil {
	// 	return "", err
	// }
	return shortURL, nil
}
