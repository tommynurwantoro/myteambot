package entity

// User _
type User struct {
	ID       uint64 `db:"id"`
	Username string `db:"username"`
	IsAdmin  bool   `db:"is_admin"`
}
