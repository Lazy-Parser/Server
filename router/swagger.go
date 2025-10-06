package router

import (
	_ "github.com/Lazy-Parser/Server/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApplySwagger(server *gin.Engine) {
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
