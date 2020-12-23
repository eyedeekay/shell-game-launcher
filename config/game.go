package config

// Game struct containers the configuration for a game
type Game struct {
	// ChrootPath is the chroot path for the game
	ChrootPath string `yaml:"ChrootPath"`
	// FileMode is the file mode to use when copying files
	FileMode string `yaml:"FileMode"`
	// Commands is the command list
	Commands []string `yaml:"Commands"`
}
