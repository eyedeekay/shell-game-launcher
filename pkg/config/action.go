package config

import (
	"fmt"
	"strings"
)

func validateAction(action string) error {
	tokens := strings.Split(action, " ")
	switch tokens[0] {
	case "chmail":
		if len(tokens) != 1 {
			return fmt.Errorf("chmail action takes no arguments")
		}
	case "login":
		if len(tokens) != 1 {
			return fmt.Errorf("login action takes no arguments")
		}
	case "menu":
		if len(tokens) != 2 {
			return fmt.Errorf("menu action takes exactly one argument")
		}
		// menu existence is tested in global config
	case "passwd":
		if len(tokens) != 1 {
			return fmt.Errorf("passwd action takes no arguments")
		}
	case "play":
		if len(tokens) != 2 {
			return fmt.Errorf("play action takes exactly one argument")
		}
		// game existence is tested in global config
	case "register":
		if len(tokens) != 1 {
			return fmt.Errorf("register action takes no arguments")
		}
	case "replay":
		if len(tokens) != 1 {
			return fmt.Errorf("replay action takes no arguments")
		}
	case "watch":
		if len(tokens) != 1 {
			return fmt.Errorf("watch action takes no arguments")
		}
	case "quit":
		if len(tokens) != 1 {
			return fmt.Errorf("quit action takes no arguments")
		}
	default:
		return fmt.Errorf("Invalid action : " + tokens[0])
	}
	return nil
}
