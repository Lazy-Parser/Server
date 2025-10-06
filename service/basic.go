package service

import "github.com/gin-gonic/gin"

// PingHandler
// @Summary Ping example
// @Description Responds with pong
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Router /basic/ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
