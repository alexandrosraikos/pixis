package handlers

import (
	"net/http"
	"strconv"

	"github.com/alexandrosraikos/pixis/database"
	"github.com/alexandrosraikos/pixis/models"
	"github.com/gin-gonic/gin"
)

// CreateDepartment handles POST /departments
func CreateDepartment(c *gin.Context) {
	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.GetDB().Create(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, department)
}

// GetDepartments handles GET /departments
func GetDepartments(c *gin.Context) {
	var departments []models.Department
	if err := database.GetDB().Find(&departments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, departments)
}

// GetDepartment handles GET /departments/:id
func GetDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}
	var department models.Department
	if err := database.GetDB().First(&department, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}
	c.JSON(http.StatusOK, department)
}

// UpdateDepartment handles PUT /departments/:id
func UpdateDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}
	var department models.Department
	db := database.GetDB()
	if err := db.First(&department, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	department.ID = uint(id)
	if err := db.Save(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, department)
}

// DeleteDepartment handles DELETE /departments/:id
func DeleteDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}
	var department models.Department
	db := database.GetDB()
	if err := db.First(&department, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}
	if err := db.Delete(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Department deleted"})
}
