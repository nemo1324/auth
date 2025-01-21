package main

import (
	"fmt"
	"log"
	"market/internal/config"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("Loaded config: %+v\n", cfg)

	// TODO: INIT LOGGER

	// TODO: инициализация приложения (app)

	// TODO: запустить grpc
}
