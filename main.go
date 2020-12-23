package main

import (
	"log"
	"shell-game-launcher/config"
)

func main() {
	config, err := config.LoadFile("example/complete.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", config)
}
