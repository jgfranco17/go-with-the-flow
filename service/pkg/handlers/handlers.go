package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jgfranco17/go-with-the-flow/core/pkg/models"
)

func VerificationHandler() func(c *gin.Context) error {
	return func(c *gin.Context) error {
		username := c.Param("user")
		c.JSON(http.StatusOK, models.VerificationMessage{
			User:    username,
			Message: "Hello, world!",
		})
		return nil
	}
}
