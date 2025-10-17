package app

import (
	"os"
	"path/filepath"
	"strconv"

	db "github.com/Lazy-Parser/Server/database"
	"github.com/Lazy-Parser/Server/database/sqlite"
	"github.com/Lazy-Parser/Server/middleware"
	"github.com/Lazy-Parser/Server/process"
	r "github.com/Lazy-Parser/Server/router"
	"github.com/gin-gonic/gin"
)

func DoWork(port int64) error {
	router := gin.Default()
	router.Use(middleware.Test()) // Use global middleware. Also, I can apply middleware for the specific endpoint

	pManager := process.NewProcessManager()
	userRepo, err := database()
	if err != nil {
		panic(err)
	}
	//// create user first
	//err = userRepo.Create(entity.User{
	//	Username: "admin",
	//})
	//if err != nil {
	//	panic(err)
	//}

	r.ApplySwagger(router)
	r.ApplyBasicRouters(router)
	r.ApplyTimerRouters(router, pManager)
	r.ApplyAuthRouters(router, userRepo)
	r.ApplyExchangeRouters(router, pManager)

	portStr := ":" + strconv.FormatInt(port, 10)
	return router.Run(portStr)
}

func database() (db.UserRepo, error) {
	wd, _ := os.Getwd()
	dbSqlite, err := sqlite.Start(filepath.Join(wd, "database/storage/app.db"), sqlite.WithAutoMigrate())
	if err != nil {
		return nil, err
	}

	return db.CreateUserRepo(dbSqlite), nil
}
