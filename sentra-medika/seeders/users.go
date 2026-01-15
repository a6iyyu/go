// Package seeders provides initial data seeding for the application.
package seeders

import (
	"fmt"
	t "time"

	m "sentra-medika/models"
	u "sentra-medika/utils"

	"gorm.io/gorm"
)

func Users(db *gorm.DB) (*m.Users, *m.Users) {
	hashedPassword, _ := u.HashPassword("admin123")

	var admin m.Users
	if err := db.Where("email = ?", "sentramedika@gmail.com").First(&admin).Error; err != nil {
		admin = m.Users{
			FullName:  "Administrator",
			Email:     "sentramedika@gmail.com",
			Password:  hashedPassword,
			Role:      string(m.Admin),
			CreatedAt: t.Now(),
		}
		db.Create(&admin)
		fmt.Println("✅ Admin created")
	}

	var doctor m.Users
	if err := db.Where("email = ?", "dr.johndoe@gmail.com").First(&doctor).Error; err != nil {
		doctor = m.Users{
			FullName:  "Dr. John Doe",
			Email:     "dr.johndoe@gmail.com",
			Password:  hashedPassword,
			Role:      string(m.Doctor),
			CreatedAt: t.Now(),
		}
		db.Create(&doctor)
		fmt.Println("✅ Doctor created")
	}

	var patient m.Users
	if err := db.Where("email = ?", "jane.smith@gmail.com").First(&patient).Error; err != nil {
		patient = m.Users{
			FullName:  "Jane Smith",
			Email:     "jane.smith@gmail.com",
			Password:  hashedPassword,
			Role:      string(m.Patient),
			CreatedAt: t.Now(),
		}
		db.Create(&patient)
		fmt.Println("✅ Patient created")
	}

	return &doctor, &patient
}