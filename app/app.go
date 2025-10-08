package app

import (
	"github.com/Lazy-Parser/Server/middleware"
	"github.com/Lazy-Parser/Server/process"
	r "github.com/Lazy-Parser/Server/router"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DoWork(port int64) error {
	router := gin.Default()
	router.Use(middleware.Test()) // Use global middleware. Also, I can apply middleware for the specific endpoint

	pManager := process.NewProcessManager()

	r.ApplySwagger(router)
	r.ApplyBasicRouters(router)
	r.ApplyTimerRouters(router, pManager)

	portStr := ":" + strconv.FormatInt(port, 10)
	return router.Run(portStr)
}
