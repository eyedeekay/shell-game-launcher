package client

import (
	"io/ioutil"
	"os"
	"reflect"
	"shell-game-launcher/pkg/config"
	"testing"
)

func TestDisplayMenu(t *testing.T) {
	realStdout := os.Stdout
	t.Cleanup(func() { os.Stdout = realStdout })
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Complete menu, while not logged in
	state := State{
		config: &config.Config{
			Menus: map[string]config.Menu{
				"test": config.Menu{
					Banner: "TEST TEST TEST",
					MenuEntries: []config.MenuEntry{
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
	want := []byte("\033[2J" +
		"TEST TEST TEST\n" +
		"\n" +
		"Not logged in.\n" +
		"\n" +
		"q) quit entry\n")
	state.displayMenu()
	// back to normal state
	w.Close()
	out, _ := ioutil.ReadAll(r)
	if !reflect.DeepEqual(out, want) {
		t.Fatalf("menu displayed incorrectly:\nwant:%+v\ngot: %+v", want, out)
	}

	// Complete menu, while logged in
	r, w, _ = os.Pipe()
	os.Stdout = w

	// Complete menu, while not logged in
	state = State{
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
		login:       "test",
	}
	want = []byte("\033[2J" +
		"TEST TEST TEST\n" +
		"\n" +
		"Logged in as: test\n" +
		"\n" +
		"w) wait entry\n" +
		"q) quit entry\n")
	state.displayMenu()
	// back to normal state
	w.Close()
	out, _ = ioutil.ReadAll(r)
	if !reflect.DeepEqual(out, want) {
		t.Fatalf("menu displayed incorrectly:\nwant:%+v\ngot: %+v", want, out)
	}
}
