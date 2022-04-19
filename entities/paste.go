package entities

import (
	"encoding/json"
	"time"
)

type Paste struct {
	Text            string    `json:"text"`
	CreatedAt       time.Time `json:"created_at"`
	ExpiryInMinutes int       `json:"expiry_in_minutes"`
}

func (c *Paste) MarshalJSON() ([]byte, error) {
	type Alias Paste
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"created_at"`
	}{
		Alias:     (*Alias)(c),
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
	})
}
