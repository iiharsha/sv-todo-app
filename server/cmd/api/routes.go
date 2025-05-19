package main

import "github.com/gin-gonic/gin"

func (app *application) routes() *gin.Engine {
	// router := gin.New()
	// router.Use(gin.Recovery())
	router := gin.Default()

	router.GET("v1/healthcheck", app.healthcheckHandler)
	router.POST("v1/todos", app.createTodoHandler)
	router.GET("v1/todos/:id", app.showTodoHandler)

	return router
}
