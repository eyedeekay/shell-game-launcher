package config

import (
	"errors"
	"regexp"
)

var reValidGameName = regexp.MustCompile(`^[\w\._]+$`)

// Game struct containers the configuration for a game
type Game struct {
	// ChrootPath is the chroot path for the game
	ChrootPath string `yaml:"ChrootPath"`
	// FileMode is the file mode to use when copying files
	FileMode string `yaml:"FileMode"`
	// Commands is the command list
	Commands []string `yaml:"Commands"`
	// ScoreFile is relative to the chroot path for the game
	ScoreCommands []string `yaml:"ScoreCommands"`
	// Env is the environment in which to exec the commands
	Env map[string]string `yaml:"Env"`
}

func (a *Game) validate(name string) error {
	// Game name
	if ok := reValidGameName.MatchString(name); !ok {
		return errors.New("Invalid Game name, must match regex `^[\\w\\._]+$` : " + name)
	}
	// ChrootPath  TODO
	// FileMode
	// Commands
	// ScoreFile
	// Env
	return nil
}
