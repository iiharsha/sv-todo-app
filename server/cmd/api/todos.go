package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iiharsha/sv-todo-app/internal/data"
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

	todo := data.Todo{
		ID:          id,
		Title:       "TestE",
		Description: "Get your test shot",
		Completed:   false,
		CreatedAt:   time.Now(),
		DueDate:     time.Now().Add(12 * time.Hour),
		Priority:    data.High,
		Tags:        []string{"gym", "testE"},
		UpdatedAt:   time.Now(),
		Version:     1,
	}

	err = app.writeJSON(c, http.StatusOK, envelope{"task": todo}, nil)
	if err != nil {
		app.logger.PrintFatal(err, nil)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "the server encountered a problem and could not process your request",
		})
	}
}
