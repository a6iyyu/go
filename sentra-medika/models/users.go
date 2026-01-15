// Package models contains the data structures and definitions for user-related entities.
package models

import (
	t "time"

	u "github.com/google/uuid"
	g "gorm.io/gorm"
)

type Role string

const (
	Admin   Role = "admin"
	Doctor  Role = "doctor"
	Patient Role = "patient"
)

type Users struct {
	ID        u.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FullName  string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"type:varchar(20);default:'patient'"`
	CreatedAt t.Time
	UpdatedAt t.Time
	DeletedAt g.DeletedAt `gorm:"index"`
}