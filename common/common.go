package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, message string, result interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"result":  result,
		"message": message,
		"success": true,
	})
}

func ResponseErr(c *gin.Context, message string, err error) {
	c.JSON(http.StatusOK, gin.H{
		"result":  err,
		"message": message,
		"success": false,
	})
}
