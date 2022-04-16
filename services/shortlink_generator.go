package services

import (
	"github.com/zaenulhilmi/pastebin/helpers"
)

type ShortlinkGenerator interface {
	Generate() (string, error)
}

func NewShortlinkGenerator(token helpers.Token) ShortlinkGenerator {
	return &md5Generator{
		token: token,
	}
}

type md5Generator struct {
	token helpers.Token
}

func (s *md5Generator) Generate() (string, error) {
	return s.token.Random(8), nil
}
