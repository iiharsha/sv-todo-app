package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) createTodoHandler(c *gin.Context) {
	fmt.Fprintln(c.Writer, "create a new movie")
}

func (app *application) showTodoHandler(c *gin.Context) {
	id, err := app.readIDParam(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
		return
	}

	c.String(http.StatusOK, "show the details of the todo %d\n", id)
}
