package handlers

import (
	"net/http"

	"github.com/alexandrosraikos/pixis/database"
	"github.com/alexandrosraikos/pixis/models"
	"github.com/gin-gonic/gin"
)

// CreateConscript handles POST /conscripts
func CreateConscript(c *gin.Context) {
	var conscript models.Conscript
	if err := c.ShouldBindJSON(&conscript); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := database.GetDB()
	if err := db.Create(&conscript).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, conscript)
}

// GetConscripts handles GET /conscripts
func GetConscripts(c *gin.Context) {
	db := database.GetDB()
	var conscripts []models.Conscript
	db.Find(&conscripts)
	c.JSON(http.StatusOK, conscripts)
}

// GetConscript handles GET /conscripts/:id
func GetConscript(c *gin.Context) {
	id := c.Param("id")
	db := database.GetDB()
	var conscript models.Conscript
	if err := db.First(&conscript, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conscript not found"})
		return
	}
	c.JSON(http.StatusOK, conscript)
}

// UpdateConscript handles PUT /conscripts/:id
func UpdateConscript(c *gin.Context) {
	id := c.Param("id")
	db := database.GetDB()
	var conscript models.Conscript
	if err := db.First(&conscript, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conscript not found"})
		return
	}
	var input models.Conscript
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&conscript).Updates(input)
	c.JSON(http.StatusOK, conscript)
}

// DeleteConscript handles DELETE /conscripts/:id
func DeleteConscript(c *gin.Context) {
	id := c.Param("id")
	db := database.GetDB()
	var conscript models.Conscript
	if err := db.First(&conscript, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conscript not found"})
		return
	}
	db.Delete(&conscript)
	c.Status(http.StatusNoContent)
}
