package config

import (
	"strings"

	"github.com/pkg/errors"
)

func validateCommand(cmd string) error {
	tokens := strings.Split(cmd, " ")
	switch tokens[0] {
	case "cp":
		if len(tokens) != 3 {
			return errors.New("cp command takes exactly two arguments")
		}
	case "exec":
		if len(tokens) <= 1 {
			return errors.New("exec command needs arguments")
		}
	case "mkdir":
		if len(tokens) != 2 {
			return errors.New("mkdir command takes exactly one argument")
		}
	case "wait":
		if len(tokens) != 1 {
			return errors.New("wait command takes no arguments")
		}
	default:
		return errors.New("Invalid command : " + tokens[0])
	}
	return nil
}
