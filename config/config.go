package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	// AppConfig is the application level configuration entries
	App App `yaml:"App"`
	// Menus is the list of menus. The first one is the default menu for an anonymous user, the second one is the default menu for an authenticated user
	Menus []Menu `yaml:"Menus"`
	// Games is the list of games.
	Games map[string]Game `yaml:"Games"`
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
	return
}
