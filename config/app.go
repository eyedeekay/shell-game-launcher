package config

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
}
