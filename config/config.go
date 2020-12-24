package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	// AppConfig is the application level configuration entries
	App App `yaml:"App"`
	// Menus is the list of menus. The first one is the default menu for an anonymous user, the second one is the default menu for an authenticated user
	Menus map[string]Menu `yaml:"Menus"`
	// Games is the list of games.
	Games map[string]Game `yaml:"Games"`
}

func (c *Config) validate() error {
	if err := c.App.validate(); err != nil {
		return err
	}
	if len(c.Menus) < 2 {
		return errors.New("A valid configuration needs at least two menu entries named anonymous and logged_in")
	}
	found_anonymous_menu := false
	found_logged_in_menu := false
	for k, v := range c.Menus {
		if err := v.validate(k); err != nil {
			return err
		}
		if k == "anonymous" {
			found_anonymous_menu = true
		}
		if k == "logged_in" {
			found_logged_in_menu = true
		}
	}
	if !found_anonymous_menu {
		return errors.New("No anonymous menu declared")
	}
	if !found_logged_in_menu {
		return errors.New("No logged_in menu declared")
	}
	for k, v := range c.Games {
		if err := v.validate(k); err != nil {
			return err
		}
	}
	// TODO menu existence is tested in global config
	// TODO game existence is tested in global config
	return nil
}

// LoadFile loads the config from a given file
func LoadFile(path string) (config Config, err error) {
	var f *os.File
	f, err = os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&config); err != nil {
		return
	}
	err = config.validate()
	return
}
