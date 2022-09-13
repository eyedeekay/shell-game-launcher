package client

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

func (s *State) Loop() error {
	for {
		s.displayMenu()
		action, err := s.getValidInput()
		if err != nil {
			log.Printf("Error: %s", err)
			continue
		}
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
			s.currentMenu = "login"
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
			if err := exec.Command(tokens[0], tokens[1:]...).Run(); err != nil {
				return err
			}
			// game existence is tested in global config
		case "register":
			if len(tokens) != 1 {
				return errors.New("register action takes no arguments")
			}
			s.currentMenu = "register"
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
			return nil
		default:
			return errors.New("Invalid action : " + tokens[0])
		}
	}
}
