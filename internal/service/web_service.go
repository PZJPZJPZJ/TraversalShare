package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func EchoText(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}
