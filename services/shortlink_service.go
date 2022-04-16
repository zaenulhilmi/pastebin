package services

import (
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/repositories"
)

type ShortlinkService interface {
	GetContent(shortlink string) (*entities.Content, error)
	CreateContent(text string, expiryInMinutes int) (string, error)
}

func NewShortlinkService(repository repositories.ShortlinkRepository, shortlinkGenerator ShortlinkGenerator) ShortlinkService {
	return &shortlinkService{
		repository: repository,
		generator:  shortlinkGenerator,
	}
}

type shortlinkService struct {
	repository repositories.ShortlinkRepository
	generator  ShortlinkGenerator
}

func (s *shortlinkService) GetContent(shortlink string) (*entities.Content, error) {
	content, err := s.repository.FindContentByShortlink(shortlink)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (s *shortlinkService) CreateContent(text string, expiryInMinutes int) (string, error) {
	shortlink := s.generator.Generate()
	err := s.repository.CreateContent(shortlink, text, expiryInMinutes)
	if err != nil {
		return "", err
	}
	return shortlink, nil
}
