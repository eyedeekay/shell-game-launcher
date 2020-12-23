package config

import (
	"reflect"
	"testing"
)

func TestLoadFile(t *testing.T) {
	_, err := LoadFile("test_data/non-existant")
	if err == nil {
		t.Fatal("non-existant config file failed without error")
	}
	_, err = LoadFile("test_data/invalid_yaml")
	if err == nil {
		t.Fatal("invalid_yaml config file failed without error")
	}
	config, err := LoadFile("../example/complete.yaml")
	want := Config{
		App: App{
			WorkingDirectory:  "var/",
			MaxUsers:          512,
			AllowRegistration: true,
			MaxNickLen:        15,
			MenuMaxIdleTime:   600,
			PostLoginCommands: []string{
				"mkdir %w/userdata/%u",
				"mkdir %w/userdata/%u/dumplog",
				"mkdir %w/userdata/%u/ttyrec",
			},
		},
		Menus: []Menu{
			Menu{
				Banner:  "Shell Game Launcher - Anonymous access%n======================================",
				XOffset: 5,
				YOffset: 2,
				MenuEntries: []MenuEntry{
					MenuEntry{
						Key:    "l",
						Label:  "login",
						Action: "login",
					},
					MenuEntry{
						Key:    "r",
						Label:  "register",
						Action: "register",
					},
					MenuEntry{
						Key:    "w",
						Label:  "watch",
						Action: "watch_menu",
					},
					MenuEntry{
						Key:    "s",
						Label:  "scores",
						Action: "scores",
					},
					MenuEntry{
						Key:    "q",
						Label:  "quit",
						Action: "quit",
					},
				},
			},
			Menu{
				Banner:  "Shell Game Launcher%n===================",
				XOffset: 5,
				YOffset: 2,
				MenuEntries: []MenuEntry{
					MenuEntry{
						Key:    "p",
						Label:  "play Nethack 3.7",
						Action: "play nethack3.7",
					},
					MenuEntry{
						Key:    "o",
						Label:  "edit game options",
						Action: "options",
					},
					MenuEntry{
						Key:    "w",
						Label:  "watch",
						Action: "watch_menu",
					},
					MenuEntry{
						Key:    "s",
						Label:  "scores",
						Action: "scores",
					},
					MenuEntry{
						Key:    "q",
						Label:  "quit",
						Action: "quit",
					},
				},
			},
		},
		Games: map[string]Game{
			"nethack3.7": Game{
				ChrootPath: "/opt/nethack",
				FileMode:   "0666",
			},
		},
	}
	if err != nil || !reflect.DeepEqual(want, config) {
		t.Fatalf("complete example failed:\nerror %v\nwant:%+v\ngot: %+v", err, want, config)
	}
}
