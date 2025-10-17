package middleware

import (
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/Lazy-Parser/Server/token"
	"github.com/gin-gonic/gin"
)

// RoleGuard - AuthGuard + role checking. Accept an acceptable list of RoleID
func RoleGuard(roles ...uint) gin.HandlerFunc {
	// do this at the initialization

	return func(c *gin.Context) {
		str := strings.Split(c.GetHeader("Authorization"), " ")
		if len(str) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		t := str[1]
		claims, err := token.ParseAccessToken(t)
		if err != nil {
			log.Println("Error token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: " + err.Error()})
			return
		}

		if !slices.Contains(roles, uint(claims.RoleID)) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
			return
		}

		c.Next()
	}
}
