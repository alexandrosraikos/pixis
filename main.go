package main

import (
	"github.com/alexandrosraikos/pixis/database"
	"github.com/alexandrosraikos/pixis/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase("database/main.db")

	r := gin.Default()

	r.POST("/conscripts", handlers.CreateConscript)
	r.GET("/conscripts", handlers.GetConscripts)
	r.GET("/conscripts/:id", handlers.GetConscript)
	r.PUT("/conscripts/:id", handlers.UpdateConscript)
	r.DELETE("/conscripts/:id", handlers.DeleteConscript)

	r.Run()
}
