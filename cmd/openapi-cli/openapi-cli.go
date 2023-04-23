package main

import (
	"log"

	"github.com/BitTraceProject/BitTrace-OpenAPI/server"
	"github.com/BitTraceProject/BitTrace-Types/pkg/config"
)

func main() {
	server.InitOpenAPIServer(config.DatabaseConfig{
		Address:  "master.collector.bittrace.proj:33062",
		Username: "openapi",
		Password: "admin",
	})

	log.Println("running openapi-cli")
	server.RunOpenAPIServer(":6060")
}
