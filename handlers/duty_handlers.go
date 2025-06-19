package handlers

import (
	"net/http"
	"strconv"

	"github.com/alexandrosraikos/pixis/database"
	"github.com/alexandrosraikos/pixis/models"
	"github.com/gin-gonic/gin"
)

// CreateDuty handles POST /duties
func CreateDuty(c *gin.Context) {
	var duty models.Duty
	if err := c.ShouldBindJSON(&duty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.GetDB().Create(&duty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, duty)
}

// GetDuties handles GET /duties
func GetDuties(c *gin.Context) {
	var duties []models.Duty
	if err := database.GetDB().Find(&duties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, duties)
}

// GetDuty handles GET /duties/:id
func GetDuty(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid duty ID"})
		return
	}
	var duty models.Duty
	if err := database.GetDB().First(&duty, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Duty not found"})
		return
	}
	c.JSON(http.StatusOK, duty)
}

// UpdateDuty handles PUT /duties/:id
func UpdateDuty(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid duty ID"})
		return
	}
	var duty models.Duty
	db := database.GetDB()
	if err := db.First(&duty, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Duty not found"})
		return
	}
	if err := c.ShouldBindJSON(&duty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	duty.ID = uint(id)
	if err := db.Save(&duty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, duty)
}

// DeleteDuty handles DELETE /duties/:id
func DeleteDuty(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid duty ID"})
		return
	}
	var duty models.Duty
	db := database.GetDB()
	if err := db.First(&duty, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Duty not found"})
		return
	}
	if err := db.Delete(&duty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Duty deleted"})
}
