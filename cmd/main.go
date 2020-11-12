package main

import (
	"log"

	RestTimeAPI "github.com/TarlexGit/time-server"
	handler "github.com/TarlexGit/time-server/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(RestTimeAPI.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while run http serv: %s", err.Error())
	}
}
