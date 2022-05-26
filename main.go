package main

import (
	"fmt"
	"log"

	"github.com/juancetusjc/twittor-back/connectiondb"
	"github.com/juancetusjc/twittor-back/handler"
)

func main() {
	fmt.Println("Init")
	if connectiondb.CheckConnection() == 0 {
		log.Fatal("Connection DB error")
	}
	handler.Handerls()

}
