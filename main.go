package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

var r = gin.Default()

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	//r.GET("/", controllers.Default)

	// Define controllers

	//	Add Endpoints

	//r.LoadHTMLGlob("templates/*.html")
	//r.Static("/static", "./static")

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	r.Use(gin.Logger())

	r.Run()
}
