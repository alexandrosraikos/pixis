package handlers

import (
	"net/http"

	"github.com/alexandrosraikos/pixis/database"
	"github.com/alexandrosraikos/pixis/models"
	"github.com/gin-gonic/gin"
)

// CreateConscript handles POST /conscripts
// @Summary Create a new conscript
// @Description Create a new conscript in the system
// @Tags conscripts
// @Accept json
// @Produce json
// @Param conscript body models.Conscript true "Conscript"
// @Success 201 {object} models.Conscript
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /conscripts [post]
func CreateConscript(c *gin.Context) {
	var conscript models.Conscript
	if err := c.ShouldBindJSON(&conscript); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	db := database.GetDB()
	if err := db.Create(&conscript).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, conscript)
}

// GetConscripts handles GET /conscripts
// @Summary List all conscripts
// @Description Get a list of all conscripts
// @Tags conscripts
// @Produce json
// @Success 200 {array} models.Conscript
// @Failure 500 {object} models.ErrorResponse
// @Router /conscripts [get]
func GetConscripts(c *gin.Context) {
	db := database.GetDB()
	var conscripts []models.Conscript
	db.Find(&conscripts)
	c.JSON(http.StatusOK, conscripts)
}

// GetConscript handles GET /conscripts/:id
// @Summary Get a conscript by ID
// @Description Get a conscript by its ID
// @Tags conscripts
// @Produce json
// @Param id path int true "Conscript ID"
// @Success 200 {object} models.Conscript
// @Failure 404 {object} models.ErrorResponse
// @Router /conscripts/{id} [get]
func GetConscript(c *gin.Context) {
	id := c.Param("id")
	db := database.GetDB()
	var conscript models.Conscript
	if err := db.First(&conscript, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Conscript not found"})
		return
	}
	c.JSON(http.StatusOK, conscript)
}

// UpdateConscript handles PUT /conscripts/:id
// @Summary Update a conscript
// @Description Update a conscript by its ID
// @Tags conscripts
// @Accept json
// @Produce json
// @Param id path int true "Conscript ID"
// @Param conscript body models.Conscript true "Conscript"
// @Success 200 {object} models.Conscript
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /conscripts/{id} [put]
func UpdateConscript(c *gin.Context) {
	id := c.Param("id")
	db := database.GetDB()
	var conscript models.Conscript
	if err := db.First(&conscript, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Conscript not found"})
		return
	}
	var input models.Conscript
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	db.Model(&conscript).Updates(input)
	c.JSON(http.StatusOK, conscript)
}

// DeleteConscript handles DELETE /conscripts/:id
// @Summary Delete a conscript
// @Description Delete a conscript by its ID
// @Tags conscripts
// @Produce json
// @Param id path int true "Conscript ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {object} models.ErrorResponse
// @Router /conscripts/{id} [delete]
func DeleteConscript(c *gin.Context) {
	id := c.Param("id")
	db := database.GetDB()
	var conscript models.Conscript
	if err := db.First(&conscript, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Conscript not found"})
		return
	}
	db.Delete(&conscript)
	c.Status(http.StatusNoContent)
}
