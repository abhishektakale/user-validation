package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LogLatency is the middleware function to log request latency
func LogLatency(c *gin.Context) {
	start := time.Now()

	// Process the request
	c.Next()

	// Log the latency asynchronously
	go func() {
		duration := time.Since(start)
		log.Printf("Request took %v", duration)
	}()
}
