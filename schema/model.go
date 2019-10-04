package schema

import "time"

// Shout defines a structure for one Shout with an id, body, and created_at
type Shout struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
