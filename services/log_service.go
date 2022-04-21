package services

import (
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/repositories"
)

type LogService interface {
	SaveLog(entities.ShortlinkLog) error
}

func NewLogService(repository repositories.LogRepository) LogService {
	return &logService{
		repository: repository,
	}
}

type logService struct {
	repository repositories.LogRepository
}

func (s *logService) SaveLog(log entities.ShortlinkLog) error {
	err := s.repository.Create(log)
	return err
}
