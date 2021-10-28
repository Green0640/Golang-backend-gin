package main

import (
	"backend/controllers"
	"backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(database.Cors())
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/fields", controllers.GetAllFields)
	router.GET("/field/:id", controllers.GetField)
	router.POST("/field", controllers.AddField)
	router.DELETE("/field/:id", controllers.DeleteField)
	router.PUT("/field/:id", controllers.UpdateField)
	router.Run(":8080")
}
