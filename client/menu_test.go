package client

import (
	"io/ioutil"
	"os"
	"reflect"
	"shell-game-launcher/config"
	"testing"
)

func TestDisplayMenu(t *testing.T) {
	realStdout := os.Stdout
	t.Cleanup(func() { os.Stdout = realStdout })
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Complete menu
	state := State{
		config: &config.Config{
			Menus: map[string]config.Menu{
				"test": config.Menu{
					Banner:  "TEST TEST TEST",
					XOffset: 5,
					YOffset: 3,
					MenuEntries: []config.MenuEntry{
						config.MenuEntry{
							Key:    "q",
							Label:  "quit",
							Action: "quit",
						},
					},
				},
			},
		},
		currentMenu: "test",
		login:       "nil",
	}
	want := []byte("\033[2J\n" +
		"TEST TEST TEST\n" +
		"\n\n\n" +
		"     q) quit\n")
	state.displayMenu()
	// back to normal state
	w.Close()
	out, _ := ioutil.ReadAll(r)
	if !reflect.DeepEqual(out, want) {
		t.Fatalf("menu displayed incorrectly:\nwant:%+v\ngot: %+v", want, out)
	}
}
