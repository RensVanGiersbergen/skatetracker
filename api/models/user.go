package models

import "time"

// User struct representing the user data in the database
type User struct {
	UserId         string     `json:"user_id" db:"user_id"`
	Username       string     `json:"username" db:"username"`
	Email          string     `json:"email" db:"email"`
	HashedPassword string     `json:"password_hash" db:"password_hash"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
}
