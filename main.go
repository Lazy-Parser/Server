package main

import (
	"log"

	"github.com/Lazy-Parser/Server/app"
)

func main() {
	log.Println("Start server...")
	if err := app.DoWork(8080); err != nil {
		panic(err)
	}
}
