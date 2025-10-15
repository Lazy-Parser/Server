package router

import (
	"github.com/Lazy-Parser/Server/database"
	"github.com/Lazy-Parser/Server/service"
	"github.com/gin-gonic/gin"
)

func ApplyAuthRouters(server *gin.Engine, userRepo database.UserRepo) {
	// use router.Group() here, and pass middleware using Options...
	server.POST("/auth/createUser", service.AuthCreateUser)
	server.POST("/auth/loginFirst", func(ctx *gin.Context) {
		service.AuthLogInFirst(ctx, userRepo)
	})
	server.POST("/auth/loginSecond", func(ctx *gin.Context) {
		service.AuthLogInSecond(ctx, userRepo)
	})
	server.GET("/auth/refresh", service.AuthRefresh)
	server.GET("/auth/logout", service.AuthLogout)
}
