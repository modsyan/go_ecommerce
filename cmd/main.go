package main

import (
	"log"

	"github.com/modsyan/go_ecommerce/cmd/api"
	"github.com/modsyan/go_ecommerce/pkg/configs"
)

func main() {
	addr := configs.Envs.Addr
	basePrefix := "/api/v1"

	server := api.NewAPIServer(addr, nil, basePrefix)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
