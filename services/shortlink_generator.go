package services

import (
	"errors"

	"github.com/zaenulhilmi/pastebin/helpers"
	"github.com/zaenulhilmi/pastebin/repositories"
)

type ShortlinkGenerator interface {
	Generate() (string, error)
}

func NewShortlinkGenerator(repository repositories.ShortlinkRepository, token helpers.Token) ShortlinkGenerator {
	return &md5Generator{
		token:      token,
		repository: repository,
	}
}

type md5Generator struct {
	token      helpers.Token
	repository repositories.ShortlinkRepository
}

func (s *md5Generator) Generate() (string, error) {
	token := s.token.Random(8)
	content, err := s.repository.FindContentByShortlink(token)
	if err != nil {
		return "", err
	}
	if content != nil {
		return "", errors.New("shortlink already exists")
	}
	return token, nil
}
