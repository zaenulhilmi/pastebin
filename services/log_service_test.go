package services_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/mocks"
	"github.com/zaenulhilmi/pastebin/services"
)

func TestSaveLog(t *testing.T) {
	logRepository := new(mocks.LogRepositoryMock)
	log := entities.ShortlinkLog{}
	logRepository.On("Create", log).Return(nil)
	logService := services.NewLogService(logRepository)
	logService.SaveLog(log)
	logRepository.AssertCalled(t, "Create", log)
}

func TestSaveLogError(t *testing.T) {
	logRepository := new(mocks.LogRepositoryMock)
	log := entities.ShortlinkLog{}
	logRepository.On("Create", log).Return(errors.New("error"))
	logService := services.NewLogService(logRepository)
	err := logService.SaveLog(log)
	assert.NotNil(t, err)
	logRepository.AssertCalled(t, "Create", log)

}
