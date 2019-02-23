package entity

import "time"

// Review _
type Review struct {
	ID        int64     `db:"id"`
	URL       string    `db:"url"`
	IsDone    bool      `db:"is_done"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
