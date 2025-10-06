package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
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
