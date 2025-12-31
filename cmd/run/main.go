package main

import (
	"log"

	"github.com/nullrish/event-collector/internal/server"
)

func main() {
	srv := server.New(":8080")
	log.Fatal(srv.Start())
}
