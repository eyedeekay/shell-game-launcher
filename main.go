package main

import (
	client "github.com/adyxax/shell-game-launcher/pkg/client"
	config "github.com/adyxax/shell-game-launcher/pkg/config"

	"log"
	"os"
)

var Configuration *config.Config
var State *client.State

func main() {
	var err error
	Configuration, err = config.LoadFile("config.yml")
	if err != nil {
		log.Printf("Error: %s", err)
		os.Exit(1)
	}
	State := client.NewState(Configuration, "")
	if err = State.Loop(); err != nil {
		log.Printf("Error: %s", err)
		os.Exit(1)
	}
}
