package main

import (
	"rinha_de_backend_vinisooo_2025/api/handler"
	"rinha_de_backend_vinisooo_2025/api/workers"
)

func main() {
	go workers.Worker()

	routes := handler.SetRoutes()
	routes.Run(":9090")
}
