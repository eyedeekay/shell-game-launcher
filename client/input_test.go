package client

import (
	"os"
	"shell-game-launcher/config"
	"testing"
)

func TestGetValidInput(t *testing.T) {
	realStdin := os.Stdin
	t.Cleanup(func() { os.Stdin = realStdin })

	// Complete menu, no input error
	state := State{
		config: &config.Config{
			Menus: map[string]config.Menu{
				"test": config.Menu{
					Banner: "TEST TEST TEST",
					MenuEntries: []config.MenuEntry{
						config.MenuEntry{
							Key:    "w",
							Label:  "wait entry",
							Action: "wait",
						},
						config.MenuEntry{
							Key:    "q",
							Label:  "quit entry",
							Action: "quit",
						},
					},
				},
			},
		},
		currentMenu: "test",
		login:       "",
	}
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Simply test quit entry
	w.WriteString("q")
	if cmd, err := state.getValidInput(); err != nil || cmd != "quit" {
		t.Fatalf("Input handled incorrectly:\nwant: wait\ngot: %s\nerror: %s\n", cmd, err)
	}
	// test quit entry after wrong keys
	w.WriteString("abcdq")
	if cmd, err := state.getValidInput(); err != nil || cmd != "quit" {
		t.Fatalf("Input handled incorrectly:\nwant: wait\ngot: %s\nerror: %s\n", cmd, err)
	}
	// test wait entry with valid quit after
	w.WriteString("wq")
	if cmd, err := state.getValidInput(); err != nil || cmd != "wait" {
		t.Fatalf("Input handled incorrectly:\nwant: wait\ngot: %s\nerror: %s\n", cmd, err)
	}
	// test input error
	w.Close()
	if cmd, err := state.getValidInput(); err == nil {
		t.Fatalf("Input handled incorrectly:\nwant: wait\ngot: %s\nerror: %s\n", cmd, err)
	}
}
