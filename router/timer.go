package router

import (
	"github.com/Lazy-Parser/Server/process"
	"github.com/Lazy-Parser/Server/service"
	"github.com/gin-gonic/gin"
)

func ApplyTimerRouters(server *gin.Engine, pm *process.Manager) {
	server.GET("/timer/start", func(c *gin.Context) {
		service.TimerStartHandler(c, pm)
	})
	server.GET("/timer/get/:id", func(c *gin.Context) {
		service.TimerGetHandler(c, pm)
	})
	server.GET("/timer/getAll", func(c *gin.Context) {
		service.TimerGetHandler(c, pm)
	})
	server.GET("/timer/stop/:id", func(c *gin.Context) {
		service.TimerStopHandler(c, pm)
	})
}
