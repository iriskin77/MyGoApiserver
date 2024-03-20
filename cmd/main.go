package main

import (
	"fmt"
	"log"

	goapiserver "github.com/iriskin77/goapiserver"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello")

	configServer := goapiserver.NewConfigServer()

	// Запуск сервера
	s := goapiserver.NewApiServer(configServer)
	if err := s.RunServer(); err != nil {
		log.Fatal(err)
	}
}
