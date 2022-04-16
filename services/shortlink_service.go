package services

import (
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/repositories"
)

type ShortlinkService interface {
	GetContent(shortlink string) (*entities.Content, error)
    CreateContent(text string, expiryInMinutes int) (string, error)
}

func NewShortlinkService(repository repositories.ShortlinkRepository) ShortlinkService {
	return &shortlinkService{
		repository: repository,
	}
}

type shortlinkService struct {
	repository repositories.ShortlinkRepository
}

func (s *shortlinkService) GetContent(shortlink string) (*entities.Content, error) {
	content, err := s.repository.FindContentByShortlink(shortlink)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (s *shortlinkService) CreateContent(text string, expiryInMinutes int) (string, error) {
    return "", nil
}
