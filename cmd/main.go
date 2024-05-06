package main

import (
	"log"

	"github.com/modsyan/go_ecommerce/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}