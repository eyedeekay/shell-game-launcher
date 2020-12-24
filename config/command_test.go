package config

import "testing"

func TestCommandValidate(t *testing.T) {
	// Empty command
	if err := validateCommand(""); err == nil {
		t.Fatal("An command cannot be empty")
	}
	// invalid command
	if err := validateCommand("invalid"); err == nil {
		t.Fatal("An command cannot be empty")
	}
	// cp
	if err := validateCommand("cp"); err == nil {
		t.Fatal("cp command needs arguments")
	}
	if err := validateCommand("cp test"); err == nil {
		t.Fatal("cp command needs exactly 2 arguments")
	}
	if err := validateCommand("cp test test test"); err == nil {
		t.Fatal("cp command needs exactly 2 arguments")
	}
	if err := validateCommand("exec test test"); err != nil {
		t.Fatal("valid exec command should be accepted")
	}
	// exec
	if err := validateCommand("exec"); err == nil {
		t.Fatal("exec command needs arguments")
	}
	if err := validateCommand("exec test"); err != nil {
		t.Fatal("valid exec command should be accepted")
	}
	// mkdir
	if err := validateCommand("mkdir"); err == nil {
		t.Fatal("mkdir command needs exactly 1 argument")
	}
	if err := validateCommand("mkdir testtest "); err == nil {
		t.Fatal("mkdir command needs exactly 1 argument")
	}
	if err := validateCommand("mkdir test"); err != nil {
		t.Fatal("valid mkdir command should be accepted")
	}
	// wait
	if err := validateCommand("wait test"); err == nil {
		t.Fatal("wait command needs no arguments")
	}
	if err := validateCommand("wait"); err != nil {
		t.Fatal("valid wait command should be accepted")
	}
}
