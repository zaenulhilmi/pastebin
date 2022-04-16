package helpers

import (
	"crypto/rand"
	"fmt"
)

type Token interface {
	Random(length int) string
}

type DefaultToken struct{}

func (t *DefaultToken) Random(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
