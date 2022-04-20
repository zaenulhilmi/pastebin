package repositories

import (
	"github.com/zaenulhilmi/pastebin/entities"
)

type LogRepository interface {
	Create(log entities.Log) error
}
