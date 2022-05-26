package main

import (
	"fmt"
	"log"

	"github.com/juancetusjc/twittor-back/handler"
)

func main() {
	fmt.Println("Init")
	if connectiondb.CheckConnection() == 0 {
		log.Fatal("Connection DBs")
	}
	handler.Handerls()

}
