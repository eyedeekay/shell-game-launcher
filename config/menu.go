package config

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var reValidMenuName = regexp.MustCompile(`^[\w\._]+$`)
var reValidKey = regexp.MustCompile(`^\w$`)

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

func (m *Menu) validate(name string) error {
	// validate name
	if ok := reValidMenuName.MatchString(name); !ok {
		return errors.New("Invalid menu name, must be an alphanumeric word and match regex `^[\\w\\._]+$` : " + name)
	}
	// Banner is just any string, nothing to validate
	// XOffset
	if m.XOffset < 0 {
		return errors.New("XOffset must be a positive integer")
	}
	// YOffset
	if m.YOffset < 0 {
		return errors.New("YOffset must be a positive integer")
	}
	// MenuEntries
	if len(m.MenuEntries) == 0 {
		return errors.New("A Menu needs MenuEntries to be valid")
	}
	// Duplicate detection is natively handled by the yaml parser
	for i := 0; i < len(m.MenuEntries); i++ {
		m.MenuEntries[i].validate()
		if m.MenuEntries[i].Action == "menu "+name {
			return errors.New("A menu shall not loop on itself")
		}
	}
	// Loop test
	return nil
}

func (m *Menu) validateConsistency(c *Config) error {
	// Necessary menus
	if _, ok := c.Menus["anonymous"]; !ok {
		return errors.New("No anonymous menu declared")
	}
	if _, ok := c.Menus["logged_in"]; !ok {
		return errors.New("No logged_in menu declared")
	}
	// Validate actions
	menus := map[string]bool{
		"anonymous": true,
		"logged_in": true,
	}
	playable := make(map[string]bool)
	for k, v := range c.Menus {
		for _, e := range v.MenuEntries {
			tokens := strings.Split(e.Action, " ")
			switch tokens[0] {
			case "menu":
				if _, ok := c.Menus[tokens[1]]; ok {
					menus[tokens[1]] = true
				} else {
					return errors.New("menu action " + tokens[1] + " in menu " + k + " does not exist")
				}
			case "play":
				if _, ok := c.Games[tokens[1]]; ok {
					playable[tokens[1]] = true
				} else {
					return errors.New("play action " + tokens[1] + " in menu " + k + " does not exist")
				}
			}
		}
	}
	// Check for unreachables
	for k, _ := range c.Menus {
		if _, ok := menus[k]; !ok {
			return errors.New("unreachable menu : " + k)
		}
	}
	for k, _ := range c.Games {
		if _, ok := playable[k]; !ok {
			return errors.New("unplayable game : " + k)
		}
	}
	return nil
}

func (m *MenuEntry) validate() error {
	// Key
	if ok := reValidKey.MatchString(m.Key); !ok {
		return errors.New("Invalid Key, must be exactly one alphanumeric character and match regex `^\\w$` : " + m.Key)
	}
	// Label
	if len(m.Label) <= 0 {
		return errors.New("Invalid Label, cannot be empty")
	}
	// Action
	if err := validateAction(m.Action); err != nil {
		return errors.Wrap(err, "Invalid Action in MenuEntry")
	}
	return nil
}
