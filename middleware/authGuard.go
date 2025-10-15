package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/Lazy-Parser/Server/token"
	"github.com/gin-gonic/gin"
)

// AuthGuard - Passes only registered users
func AuthGuard() gin.HandlerFunc {
	// do this at the initialization

	return func(c *gin.Context) {
		str := strings.Split(c.GetHeader("Authorization"), " ")
		if len(str) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		t := str[1]
		if _, err := token.ParseAccessToken(t); err != nil {
			log.Println("Error token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: " + err.Error()})
			return
		}

		c.Next()
	}
}
