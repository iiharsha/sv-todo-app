package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type envelope map[string]any

func (app *application) readIDParam(c *gin.Context) (int64, error) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) writeJSON(c *gin.Context, status int, data envelope, headers http.Header) error {
	for key, values := range headers {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	c.JSON(status, data)
	return nil
}
