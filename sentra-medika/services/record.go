// Package services provides user-related functionalities.
package services

import (
	"net/http"

	m "sentra-medika/models"
	u "sentra-medika/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateRecord(c *gin.Context) {
	var input struct {
		PatientID     string `json:"patient_id" binding:"required,uuid"`
		Diagnosis     string `json:"diagnosis" binding:"required"`
		TreatmentPlan string `json:"treatment_plan" binding:"required"`
		Notes         string `json:"notes"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctorID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	patientID, err := uuid.Parse(input.PatientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	record := m.MedicalRecords{
		DoctorID:      doctorID.(uuid.UUID),
		PatientID:     patientID,
		Diagnosis:     input.Diagnosis,
		TreatmentPlan: input.TreatmentPlan,
		Notes:         input.Notes,
	}

	if err := u.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create medical record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": record, "message": "Medical record created successfully"})
}

func GetRecords(c *gin.Context) {
	var records []m.MedicalRecords
	if err := u.DB.Preload("Doctor").Preload("Patient").Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve medical records"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": records})
}

func UpdateRecord(c *gin.Context) {
	var input struct {
		Diagnosis     string `json:"diagnosis"`
		TreatmentPlan string `json:"treatment_plan"`
		Notes         string `json:"notes"`
	}

	id := c.Param("id")
	var record m.MedicalRecords

	if err := u.DB.First(&record, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medical record not found"})
		return
	}


	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.DB.Model(&record).Updates(m.MedicalRecords{
		Diagnosis: input.Diagnosis,
		TreatmentPlan: input.TreatmentPlan,
		Notes:     input.Notes,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update medical record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": record, "message": "Medical record updated successfully"})
}

func DeleteRecord(c *gin.Context) {
	id := c.Param("id")
	var record m.MedicalRecords

	if err := u.DB.First(&record, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medical record not found"})
		return
	}

	if err := u.DB.Delete(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete medical record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Medical record deleted successfully"})
}

func GetMyRecords(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var records []m.MedicalRecords
	if err := u.DB.Where("patient_id = ?", userID).Preload("Doctor").Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve medical records"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": records})
}