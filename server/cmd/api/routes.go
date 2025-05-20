package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	// router := gin.New()
	// router.Use(gin.Recovery())
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // or "*" during development
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("v1/healthcheck", app.healthcheckHandler)
	router.POST("v1/todos", app.createTodoHandler)
	router.GET("v1/todos/:id", app.showTodoHandler)

	return router
}
