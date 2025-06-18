package main

import (
    "github.com/gin-gonic/gin"
    "pixis/database"
)

func main() {
    db := database.ConnectDatabase()
    defer db.Close()

    r := gin.Default()
    // Register your routes here

    r.Run() // default :8080
}