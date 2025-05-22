package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) logError(c *gin.Context, err error) {
	app.logger.PrintError(err, map[string]string{
		"request_method": c.Request.Method,
		"request_url":    c.Request.URL.String(),
	})
}

func (app *application) errorResponse(c *gin.Context, status int, message any) {
	env := envelope{"error": message}

	c.JSON(status, env)
}

func (app *application) serverErrorResponse(c *gin.Context, err error) {
	app.logError(c, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(c, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(c *gin.Context) {
	message := "the requested resource could not be found"

	app.errorResponse(c, http.StatusNotFound, message)
}

func (app *application) methodNotAllowed(c *gin.Context) {
	message := "method not allowed for this resource"

	app.errorResponse(c, http.StatusMethodNotAllowed, message)
}
