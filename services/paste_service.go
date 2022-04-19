package services

import (
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/repositories"
)

type PasteService interface {
	GetContent(shortlink string) (*entities.Paste, error)
	CreateContent(text string, expiryInMinutes int) (string, error)
	DeleteExpiredContent() error
}

func NewShortlinkService(readRepository repositories.ReadPasteRepository, writeRepository repositories.WritePasteRepository, shortlinkGenerator ShortlinkGenerator) PasteService {
	return &pasteService{
		readRepository:  readRepository,
		writeRepository: writeRepository,
		generator:       shortlinkGenerator,
	}
}

type pasteService struct {
	readRepository repositories.ReadPasteRepository
	writeRepository repositories.WritePasteRepository
	generator  ShortlinkGenerator
}

func (s *pasteService) GetContent(shortlink string) (*entities.Paste, error) {
	content, err := s.readRepository.FindContentByShortlink(shortlink)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (s *pasteService) CreateContent(text string, expiryInMinutes int) (string, error) {
	shortlink, err := s.generator.Generate()
	if err != nil {
		return "", err
	}
	err = s.writeRepository.CreateContent(shortlink, text, expiryInMinutes)
	if err != nil {
		return "", err
	}
	return shortlink, nil
}

func (s *pasteService) DeleteExpiredContent() error {
	err := s.writeRepository.DeleteExpiredContent()
	return err
}
