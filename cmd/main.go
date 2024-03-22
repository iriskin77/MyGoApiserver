package main

import (
	"fmt"
	"log"

	goapiserver "github.com/iriskin77/goapiserver"
	"github.com/iriskin77/goapiserver/store"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello")

	configServer := goapiserver.NewConfigServer()

	configDB := store.NewConfigDB()

	// Запуск сервера
	s := goapiserver.NewApiServer(configServer, configDB)
	if err := s.RunServer(); err != nil {
		log.Fatal(err)
	}
}
