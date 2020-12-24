package config

import "testing"

func TestGameValidate(t *testing.T) {
	empty := Game{}
	if err := empty.validate("invalid game name because of spaces"); err == nil {
		t.Fatal("invalid game name")
	}
}
