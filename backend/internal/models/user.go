package models

import (
	"time"
)

// User represents a user in the mobile app.
// The text inside the backticks `...` are GORM tags, giving specific instructions to PostgreSQL.
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"` // We will encrypt this later!
	CreatedAt time.Time
	UpdatedAt time.Time
}