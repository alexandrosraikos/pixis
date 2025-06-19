package handlers

import (
	"net/http"
	"strconv"

	"github.com/alexandrosraikos/pixis/database"
	"github.com/alexandrosraikos/pixis/models"
	"github.com/gin-gonic/gin"
)

// CreateDepartment handles POST /departments
// @Summary Create a new department
// @Description Create a new department in the system
// @Tags departments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param department body models.Department true "Department"
// @Success 201 {object} models.Department
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /departments [post]
func CreateDepartment(c *gin.Context) {
	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	if err := database.GetDB().Create(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, department)
}

// GetDepartments handles GET /departments
// @Summary List all departments
// @Description Get a list of all departments
// @Tags departments
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Department
// @Failure 500 {object} models.ErrorResponse
// @Router /departments [get]
func GetDepartments(c *gin.Context) {
	var departments []models.Department
	if err := database.GetDB().Find(&departments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, departments)
}

// GetDepartment handles GET /departments/:id
// @Summary Get a department by ID
// @Description Get a department by its ID
// @Tags departments
// @Produce json
// @Security BearerAuth
// @Param id path int true "Department ID"
// @Success 200 {object} models.Department
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /departments/{id} [get]
func GetDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid department ID"})
		return
	}
	var department models.Department
	if err := database.GetDB().First(&department, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Department not found"})
		return
	}
	c.JSON(http.StatusOK, department)
}

// UpdateDepartment handles PUT /departments/:id
// @Summary Update a department
// @Description Update a department by its ID
// @Tags departments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Department ID"
// @Param department body models.Department true "Department"
// @Success 200 {object} models.Department
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /departments/{id} [put]
func UpdateDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid department ID"})
		return
	}
	var department models.Department
	db := database.GetDB()
	if err := db.First(&department, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Department not found"})
		return
	}
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	department.ID = uint(id)
	if err := db.Save(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, department)
}

// DeleteDepartment handles DELETE /departments/:id
// @Summary Delete a department
// @Description Delete a department by its ID
// @Tags departments
// @Produce json
// @Security BearerAuth
// @Param id path int true "Department ID"
// @Success 200 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /departments/{id} [delete]
func DeleteDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid department ID"})
		return
	}
	var department models.Department
	db := database.GetDB()
	if err := db.First(&department, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Department not found"})
		return
	}
	if err := db.Delete(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.ErrorResponse{Error: "Department deleted"})
}
