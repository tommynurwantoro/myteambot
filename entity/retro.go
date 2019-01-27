package entity

import "time"

// Retro _
type Retro struct {
	ID        int64     `db:"id"`
	Username  string    `db:"username"`
	Type      string    `db:"type"`
	Message   string    `db:"message"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
