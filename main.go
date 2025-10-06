package main

import (
	"github.com/Lazy-Parser/Server/app"
	"log"
)

func main() {
	log.Println("Start server...")
	if err := app.DoWork(8080); err != nil {
		panic(err)
	}
}
