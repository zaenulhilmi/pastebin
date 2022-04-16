package services_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/pastebin/helpers"
	"github.com/zaenulhilmi/pastebin/services"
	"testing"
)

func TestMd5Generator_GenerateOk(t *testing.T) {
	generator := services.NewShortlinkGenerator(&helpers.DefaultToken{})
	shortlink, err := generator.Generate()
	assert.Nil(t, err)
	assert.NotEmpty(t, shortlink)
}
