package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Test middleware
func Test() gin.HandlerFunc {
	// do this at the initialization
	log.Println("Test middleware Init")

	return func(c *gin.Context) {
		log.Println("Test middleware request")
		c.Next()
	}
}
