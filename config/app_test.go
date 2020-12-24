package config

import (
	"os"
	"testing"
)

func TestAppvalidate(t *testing.T) {
	// WorkingDirectory
	t.Cleanup(func() { os.RemoveAll("no_permission/") })
	if err := os.Mkdir("no_permission/", 0000); err != nil {
		t.Fatal("Could not create test directory")
	}
	app := App{WorkingDirectory: "no_permission/cannot_work"}
	if err := app.validate(); err == nil {
		t.Fatal("no_permission/cannot_wor/k should not be a valid working directory")
	}
	app = App{WorkingDirectory: "no_permission/"}
	if err := app.validate(); err == nil {
		t.Fatal("no_permission/ should not be a valid working directory")
	}

	// MaxUsers
	t.Cleanup(func() { os.RemoveAll("var/") })
	app = App{
		WorkingDirectory: "var/",
		MaxUsers:         0,
	}
	if err := app.validate(); err == nil {
		t.Fatal("Negative MaxUsers should not be valid")
	}

	// AllowRegistration is just a bool, nothing to test

	// MaxNickLen
	t.Cleanup(func() { os.RemoveAll("var/") })
	app = App{
		WorkingDirectory: "var/",
		MaxUsers:         1,
		MaxNickLen:       0,
	}
	if err := app.validate(); err == nil {
		t.Fatal("Negative or zero MaxNickLen should not be valid.")
	}

	//MenuMaxIdleTime
	t.Cleanup(func() { os.RemoveAll("var/") })
	app = App{
		WorkingDirectory: "var/",
		MaxUsers:         512,
		MaxNickLen:       15,
		MenuMaxIdleTime:  0,
	}
	if err := app.validate(); err == nil {
		t.Fatal("Negative or zero MenuMaxIdleTime should not be valid.")
	}

	//PostLoginCommands is tested from command.go
}
