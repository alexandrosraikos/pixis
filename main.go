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

	// Authentication routes.
	r.POST("/auth/login", handlers.Login)

	// Protected CRUD routes.
	auth := r.Group("", handlers.AuthMiddleware())

	// Conscript CRUD routes.
	auth.POST("/conscripts", handlers.CreateConscript)
	auth.GET("/conscripts", handlers.GetConscripts)
	auth.GET("/conscripts/:id", handlers.GetConscript)
	auth.PUT("/conscripts/:id", handlers.UpdateConscript)
	auth.DELETE("/conscripts/:id", handlers.DeleteConscript)

	// Department CRUD routes.
	auth.POST("/departments", handlers.CreateDepartment)
	auth.GET("/departments", handlers.GetDepartments)
	auth.GET("/departments/:id", handlers.GetDepartment)
	auth.PUT("/departments/:id", handlers.UpdateDepartment)
	auth.DELETE("/departments/:id", handlers.DeleteDepartment)

	// Duty CRUD routes.
	auth.POST("/duties", handlers.CreateDuty)
	auth.GET("/duties", handlers.GetDuties)
	auth.GET("/duties/:id", handlers.GetDuty)
	auth.PUT("/duties/:id", handlers.UpdateDuty)
	auth.DELETE("/duties/:id", handlers.DeleteDuty)

	// Service CRUD routes.
	auth.POST("/services", handlers.CreateService)
	auth.GET("/services", handlers.GetServices)
	auth.GET("/services/:id", handlers.GetService)
	auth.PUT("/services/:id", handlers.UpdateService)
	auth.DELETE("/services/:id", handlers.DeleteService)

	// Conscript duties relationships CRUD routes.
	auth.POST("/conscript_duties", handlers.CreateConscriptDuty)
	auth.GET("/conscript_duties", handlers.GetConscriptDuties)
	auth.PUT("/conscript_duties", handlers.UpdateConscriptDuty)
	auth.DELETE("/conscript_duties", handlers.DeleteConscriptDuty)

	// Auto-generated documentation endpoints.
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
