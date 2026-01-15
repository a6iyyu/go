package services

import (
	"net/http"

	m "sentra-medika/models"
	u "sentra-medika/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var input struct {
		FullName string `json:"full_name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Role     string `json:"role" binding:"required,oneof=admin doctor patient"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := u.HashPassword(input.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := m.Users{
		FullName: input.FullName,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     input.Role,
	}

	if err := u.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user (Email might be taken)"})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"data": user, "message": "User created successfully"})
}

func GetUsers(c *gin.Context) {
	var users []m.Users

	if err := u.DB.Omit("Password").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func UpdateUser(c *gin.Context) {
	var user m.Users
	var input struct {
		FullName string `json:"full_name"`
		Email    string `json:"email" binding:"omitempty,email"` 
		Password string `json:"password" binding:"omitempty,min=6"`
		Role     string `json:"role" binding:"omitempty,oneof=admin doctor patient"`
	}

	if err := u.DB.First(&user, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateData := make(map[string]any)

	if input.FullName != "" {
		updateData["full_name"] = input.FullName
	}

	if input.Email != "" {
		updateData["email"] = input.Email
	}

	if input.Role != "" {
		updateData["role"] = input.Role
	}

	if input.Password != "" {
		hashedPassword, err := u.HashPassword(input.Password)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		updateData["password"] = hashedPassword
	}

	if err := u.DB.Model(&user).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user, "message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	var user m.Users
	id := c.Param("id")

	if err := u.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := u.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}