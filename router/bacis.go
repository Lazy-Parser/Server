// file for the bacis endpoints
package router

import (
	"github.com/Lazy-Parser/Server/middleware"
	"github.com/Lazy-Parser/Server/service"
	"github.com/gin-gonic/gin"
)

func ApplyBasicRouters(server *gin.Engine) {
	server.GET("/basic/ping", middleware.AuthGuard(), service.PingHandler)
}
