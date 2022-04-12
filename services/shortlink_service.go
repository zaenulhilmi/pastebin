package services

import (
    "github.com/zaenulhilmi/pastebin/entities"
)


type ShortlinkService interface {
    GetContent(shortlink string) (*entities.Content, error)
}


