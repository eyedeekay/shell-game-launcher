package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
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
	// App
	if err := c.App.validate(); err != nil {
		return err
	}
	// Menus
	if len(c.Menus) < 2 {
		return errors.New("A valid configuration needs at least two menu entries named anonymous and logged_in")
	}
	for k, v := range c.Menus {
		if err := v.validate(k); err != nil {
			return err
		}
	}
	// Games
	for k, v := range c.Games {
		if err := v.validate(k); err != nil {
			return err
		}
	}
	return nil
}

// LoadFile loads the c from a given file
func LoadFile(path string) (*Config, error) {
	var c *Config
	f, errOpen := os.Open(path)
	if errOpen != nil {
		return nil, errors.Wrapf(errOpen, "Failed to open configuration file %s", path)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&c); err != nil {
		return nil, errors.Wrap(err, "Failed to decode configuration file")
	}
	if err := c.validate(); err != nil {
		return nil, errors.Wrap(err, "Failed to validate configuration")
	}
	// If all looks good we validate menu consistency
	for _, v := range c.Menus {
		if err := v.validateConsistency(c); err != nil {
			return nil, errors.Wrap(err, "Failed menu consistency checks")
		}
	}
	return c, nil
}
