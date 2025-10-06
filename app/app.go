package app

import (
	"github.com/Lazy-Parser/Server/middleware"
	r "github.com/Lazy-Parser/Server/router"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DoWork(port int64) error {
	router := gin.Default()
	router.Use(middleware.Test()) // Use global middleware. Also, I can apply middleware for the specific endpoint

	r.ApplySwagger(router)
	r.ApplyBasicRouters(router)

	portStr := ":" + strconv.FormatInt(port, 10)
	return router.Run(portStr)
}
