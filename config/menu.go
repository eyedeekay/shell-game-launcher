package config

// Menu struct describes a screen menu
type Menu struct {
	// Banner is the banner to display before the menu
	Banner string `yaml:"Banner"`
	// XOffset is the X offset between the banner and the menu
	XOffset int `yaml:"XOffset"`
	// YOffset is the Y offset between the banner and the menu
	YOffset int `yaml:"YOffset"`
	// Commands is the list of commands in the menu
	MenuEntries []MenuEntry `yaml:"MenuEntries"`
}

// MenuEntry struct describes a menu entry
type MenuEntry struct {
	// Key is the key associated with the action. We need to store it as a string because of how yaml unmarshal works
	Key string `yaml:"Key"`
	// Label is the text displayed on the menu
	Label string `yaml:"Label"`
	// Action is the action executed when the menu entry is selected
	Action string `yaml:"Action"`
}
