package entities

import (
	"encoding/json"
	"time"
)

type Log struct {
	Url       string    `json:"url"`
	Method    string    `json:"method"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Log) MarshalJSON() ([]byte, error) {
	type Alias Log
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"created_at"`
	}{
		Alias:     (*Alias)(c),
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
	})
}
