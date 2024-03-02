package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/franzinBr/sin-load-balancer/config"
	"github.com/franzinBr/sin-load-balancer/internal/balancer"
	"github.com/franzinBr/sin-load-balancer/internal/server"
)

func main() {
	fmt.Println("Starting slb")

	config, err := config.LoadConfig("sin-load-balance-config-example.yml")

	if err != nil {
		log.Fatalf("Error loading configuration: %s", err.Error())
	}

	servers := server.BuildServers(config)

	b := balancer.FactoryBalancer(config, servers)

	go b.HealthCheck()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server, err := b.GetNextServer()

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		server.ServeHTTP(w, r)
	})

	log.Println("Starting slb on port", config.Sin.Port)
	err = http.ListenAndServe(config.Sin.Port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
