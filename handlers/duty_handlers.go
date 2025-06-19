package handlers

import (
	"net/http"
	"strconv"

	"github.com/alexandrosraikos/pixis/database"
	"github.com/alexandrosraikos/pixis/models"
	"github.com/gin-gonic/gin"
)

// CreateDuty handles POST /duties
// @Summary Create a new duty
// @Description Create a new duty in the system
// @Tags duties
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param duty body models.Duty true "Duty"
// @Success 201 {object} models.Duty
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /duties [post]
func CreateDuty(c *gin.Context) {
	var duty models.Duty
	if err := c.ShouldBindJSON(&duty); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	if err := database.GetDB().Create(&duty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, duty)
}

// GetDuties handles GET /duties
// @Summary List all duties
// @Description Get a list of all duties
// @Tags duties
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Duty
// @Failure 500 {object} models.ErrorResponse
// @Router /duties [get]
func GetDuties(c *gin.Context) {
	var duties []models.Duty
	if err := database.GetDB().Find(&duties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, duties)
}

// GetDuty handles GET /duties/:id
// @Summary Get a duty by ID
// @Description Get a duty by its ID
// @Tags duties
// @Produce json
// @Security BearerAuth
// @Param id path int true "Duty ID"
// @Success 200 {object} models.Duty
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /duties/{id} [get]
func GetDuty(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid duty ID"})
		return
	}
	var duty models.Duty
	if err := database.GetDB().First(&duty, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Duty not found"})
		return
	}
	c.JSON(http.StatusOK, duty)
}

// UpdateDuty handles PUT /duties/:id
// @Summary Update a duty
// @Description Update a duty by its ID
// @Tags duties
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Duty ID"
// @Param duty body models.Duty true "Duty"
// @Success 200 {object} models.Duty
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /duties/{id} [put]
func UpdateDuty(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid duty ID"})
		return
	}
	var duty models.Duty
	db := database.GetDB()
	if err := db.First(&duty, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Duty not found"})
		return
	}
	if err := c.ShouldBindJSON(&duty); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	duty.ID = uint(id)
	if err := db.Save(&duty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, duty)
}

// DeleteDuty handles DELETE /duties/:id
// @Summary Delete a duty
// @Description Delete a duty by its ID
// @Tags duties
// @Produce json
// @Security BearerAuth
// @Param id path int true "Duty ID"
// @Success 200 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /duties/{id} [delete]
func DeleteDuty(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid duty ID"})
		return
	}
	var duty models.Duty
	db := database.GetDB()
	if err := db.First(&duty, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Duty not found"})
		return
	}
	if err := db.Delete(&duty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.ErrorResponse{Error: "Duty deleted"})
}
