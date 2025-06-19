package handlers

import (
	"net/http"
	"strconv"

	"github.com/alexandrosraikos/pixis/database"
	"github.com/alexandrosraikos/pixis/models"
	"github.com/gin-gonic/gin"
)

// CreateService handles POST /services
func CreateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	if err := database.GetDB().Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, service)
}

// GetServices handles GET /services
func GetServices(c *gin.Context) {
	var services []models.Service
	if err := database.GetDB().Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

// GetService handles GET /services/:id
func GetService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid service ID"})
		return
	}
	var service models.Service
	if err := database.GetDB().First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Service not found"})
		return
	}
	c.JSON(http.StatusOK, service)
}

// UpdateService handles PUT /services/:id
func UpdateService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid service ID"})
		return
	}
	var service models.Service
	db := database.GetDB()
	if err := db.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Service not found"})
		return
	}
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	service.ID = uint(id)
	if err := db.Save(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, service)
}

// DeleteService handles DELETE /services/:id
func DeleteService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid service ID"})
		return
	}
	var service models.Service
	db := database.GetDB()
	if err := db.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Service not found"})
		return
	}
	if err := db.Delete(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.ErrorResponse{Error: "Service deleted"})
}
