package handlers

import (
	"net/http"
	"strconv"

	"github.com/alexandrosraikos/pixis/database"
	"github.com/alexandrosraikos/pixis/models"
	"github.com/gin-gonic/gin"
)

// CreateConscriptDuty assigns a duty to a conscript with metadata
func CreateConscriptDuty(c *gin.Context) {
	var cd models.ConscriptDuty
	if err := c.ShouldBindJSON(&cd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.GetDB().Create(&cd).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, cd)
}

// GetConscriptDuties lists conscript-duty assignments by conscript_id or duty_id
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cds)
}

// UpdateConscriptDuty updates metadata for a conscript-duty assignment
func UpdateConscriptDuty(c *gin.Context) {
	var input models.ConscriptDuty
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := database.GetDB()
	var cd models.ConscriptDuty
	if err := db.First(&cd, "conscript_id = ? AND duty_id = ?", input.ConscriptID, input.DutyID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assignment not found"})
		return
	}
	// Only update StartTime and EndTime if provided (zero value check)
	if !input.StartTime.IsZero() {
		cd.StartTime = input.StartTime
	}
	if !input.EndTime.IsZero() {
		cd.EndTime = input.EndTime
	}
	if err := db.Save(&cd).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cd)
}

// DeleteConscriptDuty removes a duty from a conscript
func DeleteConscriptDuty(c *gin.Context) {
	var input struct {
		ConscriptID uint `json:"conscript_id"`
		DutyID      uint `json:"duty_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := database.GetDB()
	if err := db.Delete(&models.ConscriptDuty{}, "conscript_id = ? AND duty_id = ?", input.ConscriptID, input.DutyID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
