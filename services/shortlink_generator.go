package services

import (
	"errors"

	"github.com/zaenulhilmi/pastebin/helpers"
	"github.com/zaenulhilmi/pastebin/repositories"
)

type ShortlinkGenerator interface {
	Generate() (string, error)
}

func NewShortlinkGenerator(readRepository repositories.ReadPasteRepository, token helpers.Token) ShortlinkGenerator {
	return &md5Generator{
		token:      token,
		readRepository: readRepository,
	}
}

type md5Generator struct {
	token      helpers.Token
	readRepository repositories.ReadPasteRepository
}

func (s *md5Generator) Generate() (string, error) {
	token := s.token.Random(8)
	content, err := s.readRepository.FindContentByShortlink(token)
	if err != nil {
		return "", err
	}
	if content != nil {
		return "", errors.New("shortlink already exists")
	}
	return token, nil
}
