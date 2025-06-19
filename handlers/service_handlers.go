package handlers

import (
	"net/http"
	"strconv"

	"github.com/alexandrosraikos/pixis/database"
	"github.com/alexandrosraikos/pixis/models"
	"github.com/gin-gonic/gin"
)

// CreateService handles POST /services
// @Summary Create a new service
// @Description Create a new service in the system
// @Tags services
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param service body models.Service true "Service"
// @Success 201 {object} models.Service
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /services [post]
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
// @Summary List all services
// @Description Get a list of all services
// @Tags services
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Service
// @Failure 500 {object} models.ErrorResponse
// @Router /services [get]
func GetServices(c *gin.Context) {
	var services []models.Service
	if err := database.GetDB().Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

// GetService handles GET /services/:id
// @Summary Get a service by ID
// @Description Get a service by its ID
// @Tags services
// @Produce json
// @Security BearerAuth
// @Param id path int true "Service ID"
// @Success 200 {object} models.Service
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /services/{id} [get]
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
// @Summary Update a service
// @Description Update a service by its ID
// @Tags services
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Service ID"
// @Param service body models.Service true "Service"
// @Success 200 {object} models.Service
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /services/{id} [put]
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
// @Summary Delete a service
// @Description Delete a service by its ID
// @Tags services
// @Produce json
// @Security BearerAuth
// @Param id path int true "Service ID"
// @Success 200 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /services/{id} [delete]
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
