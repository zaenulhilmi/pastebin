package entities

import (
	"encoding/json"
	"time"
)

type ShortlinkLog struct {
	Shortlink string    `json:"url"`
	Method    string    `json:"method"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *ShortlinkLog) MarshalJSON() ([]byte, error) {
	type Alias ShortlinkLog
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"created_at"`
	}{
		Alias:     (*Alias)(c),
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
	})
}
