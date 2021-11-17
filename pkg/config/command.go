package config

import (
	"fmt"
	"strings"
)

func validateCommand(cmd string) error {
	tokens := strings.Split(cmd, " ")
	switch tokens[0] {
	case "cp":
		if len(tokens) != 3 {
			return fmt.Errorf("cp command takes exactly two arguments")
		}
	case "exec":
		if len(tokens) <= 1 {
			return fmt.Errorf("exec command needs arguments")
		}
	case "mkdir":
		if len(tokens) != 2 {
			return fmt.Errorf("mkdir command takes exactly one argument")
		}
	case "wait":
		if len(tokens) != 1 {
			return fmt.Errorf("wait command takes no arguments")
		}
	default:
		return fmt.Errorf("Invalid command : " + tokens[0])
	}
	return nil
}
