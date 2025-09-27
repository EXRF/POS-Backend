package utils

import "github.com/gin-gonic/gin"

// JSONResponse standardizes API responses
func JSONResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}

// JSONError standardizes error responses
func JSONError(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, gin.H{
		"message": message,
	})
}
