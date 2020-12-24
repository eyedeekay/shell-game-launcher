package config

import "testing"

func TestActionValidate(t *testing.T) {
	// Empty action
	menuEntry := MenuEntry{Key: "l", Label: "label", Action: ""}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("An action cannot be empty")
	}
	// Invalid action
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "invalid"}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("An action must be valid")
	}
	// chmail
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "chmail a"}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("chmail action does not take arguments")
	}
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "chmail"}
	if err := menuEntry.validate(); err != nil {
		t.Fatalf("chmail action without arguments is valid\nerror: +%v", err)
	}
	// login
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "login a"}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("login action does not take arguments")
	}
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "login"}
	if err := menuEntry.validate(); err != nil {
		t.Fatalf("login action without arguments is valid\nerror: +%v", err)
	}
	// menu
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "menu"}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("menu action takes exactly one argument")
	}
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "menu test plop"}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("menu action takes exactly one argument")
	}
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "menu test"}
	if err := menuEntry.validate(); err != nil {
		t.Fatalf("menu action with one argument is valid\nerror: +%v", err)
	}
	// passwd
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "passwd a"}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("passwd action does not take arguments")
	}
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "passwd"}
	if err := menuEntry.validate(); err != nil {
		t.Fatalf("passwd action without arguments is valid\nerror: +%v", err)
	}
	// play
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "play"}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("play action takes exactly one argument")
	}
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "play test plop"}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("play action takes exactly one argument")
	}
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "play test"}
	if err := menuEntry.validate(); err != nil {
		t.Fatalf("play action with one argument is valid\nerror: +%v", err)
	}
	// register
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "register a"}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("register action does not take arguments")
	}
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "register"}
	if err := menuEntry.validate(); err != nil {
		t.Fatalf("register action without arguments is valid\nerror: +%v", err)
	}
	// replay
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "replay a"}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("replay action does not take arguments")
	}
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "replay"}
	if err := menuEntry.validate(); err != nil {
		t.Fatalf("replay action without arguments is valid\nerror: +%v", err)
	}
	// watch
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "watch a"}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("watch action does not take arguments")
	}
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "watch"}
	if err := menuEntry.validate(); err != nil {
		t.Fatalf("watch action without arguments is valid\nerror: +%v", err)
	}
	// quit
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "quit a"}
	if err := menuEntry.validate(); err == nil {
		t.Fatal("quit action does not take arguments")
	}
	menuEntry = MenuEntry{Key: "l", Label: "label", Action: "quit"}
	if err := menuEntry.validate(); err != nil {
		t.Fatalf("quit action without arguments is valid\nerror: +%v", err)
	}
}
