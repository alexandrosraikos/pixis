package main

import (
	"github.com/alexandrosraikos/pixis/database"
	"github.com/alexandrosraikos/pixis/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// This blank import is required for swaggo/swag to serve the generated docs.
	_ "github.com/alexandrosraikos/pixis/docs"
)

func main() {
	database.ConnectDatabase("database/main.db")

	r := gin.Default()

	// Conscript CRUD routes.
	r.POST("/conscripts", handlers.CreateConscript)
	r.GET("/conscripts", handlers.GetConscripts)
	r.GET("/conscripts/:id", handlers.GetConscript)
	r.PUT("/conscripts/:id", handlers.UpdateConscript)
	r.DELETE("/conscripts/:id", handlers.DeleteConscript)

	// Department CRUD routes.
	r.POST("/departments", handlers.CreateDepartment)
	r.GET("/departments", handlers.GetDepartments)
	r.GET("/departments/:id", handlers.GetDepartment)
	r.PUT("/departments/:id", handlers.UpdateDepartment)
	r.DELETE("/departments/:id", handlers.DeleteDepartment)

	// Duty CRUD routes.
	r.POST("/duties", handlers.CreateDuty)
	r.GET("/duties", handlers.GetDuties)
	r.GET("/duties/:id", handlers.GetDuty)
	r.PUT("/duties/:id", handlers.UpdateDuty)
	r.DELETE("/duties/:id", handlers.DeleteDuty)

	// Service CRUD routes.
	r.POST("/services", handlers.CreateService)
	r.GET("/services", handlers.GetServices)
	r.GET("/services/:id", handlers.GetService)
	r.PUT("/services/:id", handlers.UpdateService)
	r.DELETE("/services/:id", handlers.DeleteService)

	// Conscript duties relationships CRUD routes.
	r.POST("/conscript_duties", handlers.CreateConscriptDuty)
	r.GET("/conscript_duties", handlers.GetConscriptDuties)
	r.PUT("/conscript_duties", handlers.UpdateConscriptDuty)
	r.DELETE("/conscript_duties", handlers.DeleteConscriptDuty)

	// Auto-generated documentation endpoints.
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
