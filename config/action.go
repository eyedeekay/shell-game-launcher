package config

import (
	"strings"

	"github.com/pkg/errors"
)

func validateAction(action string) error {
	tokens := strings.Split(action, " ")
	switch tokens[0] {
	case "chmail":
		if len(tokens) != 1 {
			return errors.New("chmail action takes no arguments")
		}
	case "login":
		if len(tokens) != 1 {
			return errors.New("login action takes no arguments")
		}
	case "menu":
		if len(tokens) != 2 {
			return errors.New("menu action takes exactly one argument")
		}
		// menu existence is tested in global config
	case "passwd":
		if len(tokens) != 1 {
			return errors.New("passwd action takes no arguments")
		}
	case "play":
		if len(tokens) != 2 {
			return errors.New("play action takes exactly one argument")
		}
		// game existence is tested in global config
	case "register":
		if len(tokens) != 1 {
			return errors.New("register action takes no arguments")
		}
	case "replay":
		if len(tokens) != 1 {
			return errors.New("replay action takes no arguments")
		}
	case "watch":
		if len(tokens) != 1 {
			return errors.New("watch action takes no arguments")
		}
	case "quit":
		if len(tokens) != 1 {
			return errors.New("quit action takes no arguments")
		}
	default:
		return errors.New("Invalid action : " + tokens[0])
	}
	return nil
}
