package config

import (
	"regexp"

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
	if m.XOffset <= 0 {
		return errors.New("XOffset must be a positive integer")
	}
	// YOffset
	if m.YOffset <= 0 {
		return errors.New("YOffset must be a positive integer")
	}
	// MenuEntries
	keys := make(map[string]bool)
	for i := 0; i < len(m.MenuEntries); i++ {
		m.MenuEntries[i].validate()
		if _, duplicate := keys[m.MenuEntries[i].Key]; duplicate {
			return errors.New("A Menu has a duplicate key " + m.MenuEntries[i].Key)
		}
		keys[m.MenuEntries[i].Key] = true
		if m.MenuEntries[i].Action == "menu "+name {
			return errors.New("A menu shall not loop on itself")
		}
	}
	// Loop test
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
