package config

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadFile(t *testing.T) {
	// Non existant file
	_, err := LoadFile("test_data/non-existant")
	if err == nil {
		t.Fatal("non-existant config file failed without error")
	}

	// Invalid yaml file
	_, err = LoadFile("test_data/invalid_yaml")
	if err == nil {
		t.Fatal("invalid_yaml config file failed without error")
	}

	// Minimal yaml file
	_, err = LoadFile("test_data/minimal.yaml")
	want := Config{
		App: App{
			WorkingDirectory:  "var/",
			MaxUsers:          1,
			AllowRegistration: true,
			MaxNickLen:        15,
			MenuMaxIdleTime:   600,
		},
		Menus: map[string]Menu{
			"anonymous": Menu{
				MenuEntries: []MenuEntry{
					MenuEntry{
						Key:    "q",
						Label:  "quit",
						Action: "quit",
					},
				},
			},
			"logged_in": Menu{
				MenuEntries: []MenuEntry{
					MenuEntry{
						Key:    "q",
						Label:  "quit",
						Action: "quit",
					},
				},
			},
		},
	}
	if config, err := LoadFile("test_data/minimal.yaml"); err != nil || !reflect.DeepEqual(want, config) {
		t.Fatalf("minimal example failed:\nerror %v\nwant:%+v\ngot: %+v", err, want, config)
	}

	// TODO test non existant game in play actions, and duplicate
	//menuEntry = MenuEntry{
	//Key:    "p",
	//Label:  "play non existant game",
	//Action: "play nonexistant",
	//}
	//if err := menuEntry.validate(); err == nil {
	//t.Fatal("An inexistant game cannot be played")
	//}

	t.Cleanup(func() { os.RemoveAll("var/") })
	// Invalid App
	if _, err := LoadFile("test_data/invalid_app.yaml"); err == nil {
		t.Fatal("Invalid App entry should fail to load")
	}
	// Not enough menus
	if _, err := LoadFile("test_data/not_enough_menus.yaml"); err == nil {
		t.Fatal("not enough menu entries should fail to load")
	}
	// Invalid Menus
	if _, err := LoadFile("test_data/invalid_menus.yaml"); err == nil {
		t.Fatal("Invalid menu entry should fail to load")
	}
	// no anonymous Menu
	if _, err := LoadFile("test_data/no_anonymous_menu.yaml"); err == nil {
		t.Fatal("Invalid menu entry should fail to load")
	}
	// no logged_in Menu
	if _, err := LoadFile("test_data/no_logged_in_menu.yaml"); err == nil {
		t.Fatal("Invalid menu entry should fail to load")
	}
	// duplicate menu
	if _, err := LoadFile("test_data/duplicate_menu.yaml"); err == nil {
		t.Fatal("duplicate menu should fail to load")
	}
	// non existant menu action referenced
	if _, err := LoadFile("test_data/non_existant_menu.yaml"); err == nil {
		t.Fatal("menu entry referencing a non existant menu should fail to load")
	}
	// non existant game referenced in play action
	if _, err := LoadFile("test_data/non_existant_game.yaml"); err == nil {
		t.Fatal("menu entry referencing a non existant play action should fail to load")
	}
	// unreachable menu
	if _, err := LoadFile("test_data/unreachable_menu.yaml"); err == nil {
		t.Fatal("unreachable menu should fail to load")
	}
	// unreachable game
	if _, err := LoadFile("test_data/unreachable_game.yaml"); err == nil {
		t.Fatal("unreachable game should fail to load")
	}

	// Complexe example
	want = Config{
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
		Menus: map[string]Menu{
			"anonymous": Menu{
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
						Key:    "q",
						Label:  "quit",
						Action: "quit",
					},
				},
			},
			"logged_in": Menu{
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
						Action: "menu options",
					},
					MenuEntry{
						Key:    "w",
						Label:  "watch",
						Action: "watch",
					},
					MenuEntry{
						Key:    "r",
						Label:  "replay",
						Action: "replay",
					},
					MenuEntry{
						Key:    "c",
						Label:  "change password",
						Action: "passwd",
					},
					MenuEntry{
						Key:    "m",
						Label:  "change email",
						Action: "chmail",
					},
					MenuEntry{
						Key:    "q",
						Label:  "quit",
						Action: "quit",
					},
				},
			},
			"options": Menu{
				Banner:  "Options%n=======",
				XOffset: 5,
				YOffset: 2,
				MenuEntries: []MenuEntry{
					MenuEntry{
						Key:    "z",
						Label:  "back",
						Action: "menu logged_in",
					},
				},
			},
		},
		Games: map[string]Game{
			"nethack3.7": Game{
				ChrootPath: "/opt/nethack",
				FileMode:   "0666",
				ScoreCommands: []string{
					"exec /games/nethack -s all",
					"wait",
				},
				Commands: []string{
					"cp /games/var/save/%u%n.gz /games/var/save/%u%n.gz.bak",
					"exec /games/nethack -u %n",
				},
				Env: map[string]string{
					"NETHACKOPTIONS": "@%ruserdata/%n/%n.nhrc",
				},
			},
		},
	}
	if config, err := LoadFile("../example/complete.yaml"); err != nil || !reflect.DeepEqual(want, config) {
		t.Fatalf("complete example failed:\nerror %v\nwant:%+v\ngot: %+v", err, want, config)
	}
}
