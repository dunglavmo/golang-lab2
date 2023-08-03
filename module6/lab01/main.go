package main

import (
	"fmt"
	"log"
	"net/http"

	"loadbalancer/loadbalancer"
	"loadbalancer/middleware"
	"loadbalancer/repositories"
)

func main() {
	config, err := loadbalancer.ReadConfig()
	if err != nil {
		log.Fatal("Error reading system.conf:", err)
	}

	repo, err := repositories.NewRedisRepository("localhost:6379")
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	lb := loadbalancer.NewLoadBalancer(config, repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/", lb.Handler)

	if config.LoggingEnabled {
		muxWithLogger := middleware.Logger(mux)
		http.Handle("/", muxWithLogger)
	} else {
		http.Handle("/", mux)
	}

	port := fmt.Sprintf(":%d", config.Port)
	fmt.Printf("Load Balancer listening on port %s...\n", port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal("Error starting the load balancer:", err)
	}
}
