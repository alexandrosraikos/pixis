package handlers

import (
	"net/http"
	"strconv"

	"github.com/alexandrosraikos/pixis/database"
	"github.com/alexandrosraikos/pixis/models"
	"github.com/gin-gonic/gin"
)

// CreateConscriptDuty assigns a duty to a conscript with metadata
// @Summary Assign a duty to a conscript
// @Description Assign a duty to a conscript with start and end time
// @Tags conscript_duties
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param conscript_duty body models.ConscriptDuty true "ConscriptDuty"
// @Success 201 {object} models.ConscriptDuty
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /conscript_duties [post]
func CreateConscriptDuty(c *gin.Context) {
	var cd models.ConscriptDuty
	if err := c.ShouldBindJSON(&cd); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	if err := database.GetDB().Create(&cd).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, cd)
}

// GetConscriptDuties lists conscript-duty assignments by conscript_id or duty_id
// @Summary List conscript-duty assignments
// @Description List conscript-duty assignments by conscript_id or duty_id
// @Tags conscript_duties
// @Produce json
// @Security BearerAuth
// @Param conscript_id query int false "Conscript ID"
// @Param duty_id query int false "Duty ID"
// @Success 200 {array} models.ConscriptDuty
// @Failure 500 {object} models.ErrorResponse
// @Router /conscript_duties [get]
func GetConscriptDuties(c *gin.Context) {
	var cds []models.ConscriptDuty
	db := database.GetDB()
	conscriptID := c.Query("conscript_id")
	dutyID := c.Query("duty_id")

	query := db
	if conscriptID != "" {
		id, err := strconv.Atoi(conscriptID)
		if err == nil {
			query = query.Where("conscript_id = ?", id)
		}
	}
	if dutyID != "" {
		id, err := strconv.Atoi(dutyID)
		if err == nil {
			query = query.Where("duty_id = ?", id)
		}
	}
	if err := query.Find(&cds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, cds)
}

// UpdateConscriptDuty updates metadata for a conscript-duty assignment
// @Summary Update a conscript-duty assignment
// @Description Update start and end time for a conscript-duty assignment
// @Tags conscript_duties
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param conscript_duty body models.ConscriptDuty true "ConscriptDuty"
// @Success 200 {object} models.ConscriptDuty
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /conscript_duties [put]
func UpdateConscriptDuty(c *gin.Context) {
	var input models.ConscriptDuty
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	db := database.GetDB()
	var cd models.ConscriptDuty
	if err := db.First(&cd, "conscript_id = ? AND duty_id = ?", input.ConscriptID, input.DutyID).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Assignment not found"})
		return
	}
	if !input.StartTime.IsZero() {
		cd.StartTime = input.StartTime
	}
	if !input.EndTime.IsZero() {
		cd.EndTime = input.EndTime
	}
	if err := db.Save(&cd).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, cd)
}

// DeleteConscriptDuty removes a duty from a conscript
// @Summary Remove a duty from a conscript
// @Description Remove a duty from a conscript by conscript_id and duty_id
// @Tags conscript_duties
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param conscript_duty body object true "ConscriptDuty IDs" {"conscript_id":0,"duty_id":0}
// @Success 204 {string} string "No Content"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /conscript_duties [delete]
func DeleteConscriptDuty(c *gin.Context) {
	var input struct {
		ConscriptID uint `json:"conscript_id"`
		DutyID      uint `json:"duty_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	db := database.GetDB()
	if err := db.Delete(&models.ConscriptDuty{}, "conscript_id = ? AND duty_id = ?", input.ConscriptID, input.DutyID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
