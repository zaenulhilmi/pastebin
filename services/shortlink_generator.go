package services

import (
	"crypto/rand"
	"fmt"
)

type ShortlinkGenerator interface {
	Generate() string
}

func NewShortlinkGenerator() ShortlinkGenerator {
	return &md5Generator{}
}

type emptyStringGenerator struct{}

func (s *emptyStringGenerator) Generate() string {
	return ""
}

type md5Generator struct{}

func (s *md5Generator) Generate() string {
	return randToken()
}

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
