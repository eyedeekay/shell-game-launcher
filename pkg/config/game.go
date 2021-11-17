package config

import (
	"regexp"

	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

var reValidGameName = regexp.MustCompile(`^[\w\._]+$`)
var reValidFileMode = regexp.MustCompile(`^0[\d]{3}$`)
var reSpace = regexp.MustCompile(`^\s$`)

// Game struct containers the configuration for a game
type Game struct {
	// ChrootPath is the chroot path for the game
	ChrootPath string `yaml:"ChrootPath"`
	// FileMode is the file mode to use when copying files
	FileMode string `yaml:"FileMode"`
	// Commands is the command list
	Commands []string `yaml:"Commands"`
	// Env is the environment in which to exec the commands
	Env map[string]string `yaml:"Env"`
}

func (g *Game) validate(name string) error {
	// Game name
	if ok := reValidGameName.MatchString(name); !ok {
		return errors.New("Invalid Game name, must match regex `^[\\w\\._]+$` : " + name)
	}
	// ChrootPath  TODO
	if err := unix.Access(g.ChrootPath, unix.R_OK|unix.X_OK); err != nil {
		return errors.Wrapf(err, "Invalid ChrootPath : %s", g.ChrootPath)
	}
	// FileMode
	if ok := reValidFileMode.MatchString(g.FileMode); !ok {
		return errors.New("Invalid File Mode, must match regex `^0[\\d]{3}$` : " + name)
	}
	// Commands
	if len(g.Commands) == 0 {
		return errors.New("Invalid game " + name + " has no commands")
	}
	for i := 0; i < len(g.Commands); i++ {
		if err := validateCommand(g.Commands[i]); err != nil {
			return errors.Wrapf(err, "Failed to validate Commands for game %s", name)
		}
	}
	// Env
	for k, _ := range g.Env {
		for _, c := range k {
			switch c {
			case '=':
				return errors.New("Environment variable key must not contain equal sign")
			case '\000':
				return errors.New("Environment variable key must not contain null character")
			}
			if reSpace.MatchString(string(c)) {
				return errors.New("Environment variable key must not contain spaces")
			}
		}
	}
	return nil
}
