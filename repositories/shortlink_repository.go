package repositories

import (
	"github.com/zaenulhilmi/pastebin/entities"
)

type ShortlinkRepository interface {
	FindContentByShortlink(shortlink string) (*entities.Content, error)
}
