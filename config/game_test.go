package config

import "testing"

func TestGameValidate(t *testing.T) {
	// Game name
	game := Game{}
	if err := game.validate("invalid game name because of spaces"); err == nil {
		t.Fatal("game name with spaces should not be valid")
	}
	// ChrootPath
	game = Game{ChrootPath: "test_data/non_existant"}
	if err := game.validate("test"); err == nil {
		t.Fatal("non_existant ChrootPath should not be valid")
	}
	// FileMode
	game = Game{
		ChrootPath: "test_data/fake_nethack_directory",
	}
	if err := game.validate("test"); err == nil {
		t.Fatal("Invalid FileMode should not be valid")
	}
	game = Game{
		ChrootPath: "test_data/fake_nethack_directory",
		FileMode:   "abcd",
	}
	if err := game.validate("test"); err == nil {
		t.Fatal("Invalid FileMode should not be valid")
	}
	game = Game{
		ChrootPath: "test_data/fake_nethack_directory",
		FileMode:   "777",
	}
	if err := game.validate("test"); err == nil {
		t.Fatal("Invalid FileMode should not be valid")
	}
	// Commands are mostly tested from command_test.go
	game = Game{
		ChrootPath: "test_data/fake_nethack_directory",
		FileMode:   "0777",
	}
	if err := game.validate("test"); err == nil {
		t.Fatal("Empty Commands list should not be valid")
	}
	game = Game{
		ChrootPath: "test_data/fake_nethack_directory",
		FileMode:   "0777",
		Commands:   []string{"invalid"},
	}
	if err := game.validate("test"); err == nil {
		t.Fatal("Invalid command in Commands should not be valid")
	}
	// Env
	game = Game{
		ChrootPath: "test_data/fake_nethack_directory",
		FileMode:   "0777",
		Commands:   []string{"wait"},
	}
	if err := game.validate("test"); err != nil {
		t.Fatal("Empty env list should be valid")
	}
	game = Game{
		ChrootPath: "test_data/fake_nethack_directory",
		FileMode:   "0777",
		Commands:   []string{"wait"},
		Env: map[string]string{
			"test invalid": "test",
		},
	}
	if err := game.validate("test"); err == nil {
		t.Fatal("Spaces in environnement variable name are invalid")
	}
	game = Game{
		ChrootPath: "test_data/fake_nethack_directory",
		FileMode:   "0777",
		Commands:   []string{"wait"},
		Env: map[string]string{
			"test\000invalid": "test",
		},
	}
	if err := game.validate("test"); err == nil {
		t.Fatal("null character in environnement variable name are invalid")
	}
	game = Game{
		ChrootPath: "test_data/fake_nethack_directory",
		FileMode:   "0777",
		Commands:   []string{"wait"},
		Env: map[string]string{
			"test=invalid": "test",
		},
	}
	if err := game.validate("test"); err == nil {
		t.Fatal("equals symbol in environnement variable name are invalid")
	}
	// Valid Game entry
	game = Game{
		ChrootPath: "test_data/fake_nethack_directory",
		FileMode:   "0777",
		Commands:   []string{"wait"},
		Env: map[string]string{
			"test": "test",
		},
	}
	if err := game.validate("test"); err != nil {
		t.Fatalf("Valid game entry should pass but got error %s", err)
	}
}
