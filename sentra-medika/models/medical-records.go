// Package models contains the data models for the application.
package models

import (
	t "time"

	u "github.com/google/uuid"
	g "gorm.io/gorm"
)

type MedicalRecords struct {
	ID            u.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PatientID     u.UUID      `gorm:"type:uuid;not null;index"`
	Patient       Users       `gorm:"foreignKey:PatientID"`
	DoctorID      u.UUID      `gorm:"type:uuid;not null;index"`
	Doctor        Users       `gorm:"foreignKey:DoctorID"`
	Diagnosis     string      `gorm:"type:varchar(255);not null"`
	TreatmentPlan string      `gorm:"type:text;not null"`
	Notes         string      `gorm:"type:text"`
	CreatedAt     t.Time      `gorm:"autoCreateTime"`
	UpdatedAt     t.Time      `gorm:"autoUpdateTime"`
	DeletedAt     g.DeletedAt `gorm:"index"`
}