// Package seeders is used to seed initial data into the database.
package seeders

import (
	"fmt"

	m "sentra-medika/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func MedicalRecords(db *gorm.DB, doctorID uuid.UUID, patientID uuid.UUID) {
	if doctorID == uuid.Nil || patientID == uuid.Nil {
		fmt.Println("❌ Seeding failed because the Doctor or Patient ID is invalid.")
		return
	}

	records := []m.MedicalRecords{
		{
			PatientID:     patientID,
			DoctorID:      doctorID,
			Diagnosis:     "Flu Ringan",
			TreatmentPlan: "Istirahat total dan minum vitamin C",
			Notes:         "Annual physical examination. All vitals are normal.",
		},
		{
			PatientID:     patientID,
			DoctorID:      doctorID,
			Diagnosis:     "Alergi Debu",
			TreatmentPlan: "Cetirizine 10mg",
			Notes:         "Treated for seasonal allergies. Prescribed antihistamines.",
		},
	}

	if err := db.Create(&records).Error; err != nil {
		fmt.Printf("❌ Failed to seed medical records: %v\n", err)
	} else {
		fmt.Println("✅ Medical Records seeded successfully")
	}
}