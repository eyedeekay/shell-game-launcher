package config

import "testing"

func TestMenuValidate(t *testing.T) {
	// menu name
	menu := Menu{}
	if err := menu.validate(""); err == nil {
		t.Fatal("Empty menu name is not valid")
	}
	if err := menu.validate("test test"); err == nil {
		t.Fatal("non alphanumeric menu name is not valid")
	}
	// Banner is just any string, nothing to validate
	// XOffset
	menu = Menu{XOffset: -1}
	if err := menu.validate("test"); err == nil {
		t.Fatal("Negative XOffset should not be valid")
	}
	// YOffset
	menu = Menu{
		XOffset: 1,
		YOffset: -1,
	}
	if err := menu.validate("test"); err == nil {
		t.Fatal("Negative YOffset should not be valid")
	}
	// MenuEntries are mostly tested bellow
	menu = Menu{}
	if err := menu.validate("test"); err == nil {
		t.Fatal("A menu without menu entries should not be valid")
	}
	// loop menu
	menu = Menu{
		XOffset: 1,
		YOffset: 1,
		MenuEntries: []MenuEntry{
			MenuEntry{
				Key:    "a",
				Label:  "test",
				Action: "menu test",
			},
		},
	}
	if err := menu.validate("test"); err == nil {
		t.Fatal("A menu should not be able to loop on itself")
	}
	// A valid menu
	menu = Menu{
		XOffset: 1,
		YOffset: 1,
		MenuEntries: []MenuEntry{
			MenuEntry{
				Key:    "a",
				Label:  "test",
				Action: "quit",
			},
		},
	}
	if err := menu.validate("test"); err != nil {
		t.Fatal("A valid menu should pass")
	}
}

func TestMenuEntryValidate(t *testing.T) {
	// Key
	menuEntry := MenuEntry{}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("A Key cannot be empty")
	}
	menuEntry = MenuEntry{Key: "ab"}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("A Key should be only one character")
	}
	menuEntry = MenuEntry{Key: " "}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("A Key should be a printable character")
	}
	// Label
	menuEntry = MenuEntry{
		Key:   "l",
		Label: "",
	}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("A Label cannot be empty")
	}
	// Actions are tested in action_test.go
}
