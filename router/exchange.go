package router

import (
	"github.com/Lazy-Parser/Server/process"
	"github.com/Lazy-Parser/Server/service"
	"github.com/gin-gonic/gin"
)

func ApplyExchangeRouters(server *gin.Engine, pManager *process.Manager) {
	//roles := [1]int{} // just admin

	// do not forget to pass roleIDs into middleware.RoleGuard()
	server.GET("/ex/list", service.ExList)
	server.GET("/ex/:name/start", func(ctx *gin.Context) {
		service.ExStart(ctx, pManager)
	})
}
