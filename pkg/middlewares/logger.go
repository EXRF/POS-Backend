package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// LoggerToFile creates a logger middleware that writes logs in a custom format
func LoggerToFile() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN] %s | %3d | %13v | %15s | %-7s %#v\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
		)
	})
}

// RecoveryWithLogger recovers from panics and logs the error with stack trace
func RecoveryWithLogger() gin.HandlerFunc {
	// Use os.Stderr (or os.Stdout) since gin expects an io.Writer
	return gin.CustomRecoveryWithWriter(os.Stderr, func(c *gin.Context, recovered interface{}) {
		switch err := recovered.(type) {
		case string:
			log.Printf("panic occurred: %s\n%s", err, debug.Stack())
		case error:
			log.Printf("panic occurred: %v\n%s", err, debug.Stack())
		default:
			log.Printf("panic occurred: %#v\n%s", err, debug.Stack())
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
	})
}
