package service

import (
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

func (svc *URLService) ShortenURL(longURL string) (string, error) {
	shortURL := utils.GenerateShortURL()
	// err := svc.URLRepo.StoreURL(shortURL, longURL)
	// if err != nil {
	// 	return "", err
	// }
	return shortURL, nil
}

func (svc *URLService) GetLongURL(shortURL string) (string, error) {
	longURL := " -------- "
	// longURL, err := svc.URLRepo.FetchURL(shortURL)
	// if err != nil {
	// 	return "", err
	// }
	return longURL, nil
}

func (svc *URLService) DeleteURL(shortURL string) error {
	// err := svc.URLRepo.DeleteURL(shortURL)
	// if err != nil {
	// 	return err
	// }
	return nil
}
