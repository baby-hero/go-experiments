package main

import (
	"fmt"
	"go-experiments/config"
	"log"
)

func main() {
	name := "Go Developers"
	fmt.Println("Hello for", name)
	cfg, err := config.LoadConfig("config/config.yaml")

	if err != nil {
		log.Fatalf("Failed to load config %v", err)
	}
	fmt.Printf("Database config: %+v\n", cfg.Database)
}
