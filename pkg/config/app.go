package config

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

// App struct contains the configuration for this application
type App struct {
	// WorkingDirectory is the program working directory where the user data, save files and scores are stored
	WorkingDirectory string `yaml:"WorkingDirectory"`
	// MaxUsers is the maximum amount of registered users to allow
	MaxUsers int `yaml:"MaxUsers"`
	// AllowRegistration allows registration of new users
	AllowRegistration bool `yaml:"AllowRegistration"`
	// MaxNickLen Maximum length for a nickname
	MaxNickLen int `yaml:"MaxNickLen"`
	// MenuMaxIdleTime is the maximum number of seconds a user can be idle on the menu before the program exits
	MenuMaxIdleTime int `yaml:"MenuMaxIdleTime"`
	// PostLoginCommands is the list of commands to execute upon login, like creating save directories for games
	PostLoginCommands []string `yaml:"PostLoginCommands"`
	// TODO admin contact
}

func (a *App) validate() error {
	// WorkingDirectory
	if err := os.MkdirAll(a.WorkingDirectory, 0700); err != nil {
		return fmt.Errorf("Invalid WorkingDirectory %s : %w", a.WorkingDirectory, err)
	}
	if err := unix.Access(a.WorkingDirectory, unix.W_OK|unix.R_OK|unix.X_OK); err != nil {
		return fmt.Errorf("invalid WorkingDirectory %s : %w", a.WorkingDirectory, err)
	}
	// MaxUsers
	if a.MaxUsers <= 0 {
		return fmt.Errorf("MaxUsers must be a positive integer")
	}
	// AllowRegistration is just a bool, nothing to validate
	// MaxNickLen
	if a.MaxNickLen <= 0 {
		return fmt.Errorf("MaxNickLen must be a positive integer")
	}
	// MenuMaxIdleTime
	if a.MenuMaxIdleTime <= 0 {
		return fmt.Errorf("MenuMaxIdleTime must be a positive integer")
	}
	// PostLoginCommands
	for i := 0; i < len(a.PostLoginCommands); i++ {
		if err := validateCommand(a.PostLoginCommands[i]); err != nil {
			return fmt.Errorf("Failed to validate PostLoginCommands : %w", err)
		}
	}
	return nil
}
