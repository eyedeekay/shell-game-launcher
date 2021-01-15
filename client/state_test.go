package client

import "testing"

func TestNewState(t *testing.T) {
	// Empty login
	if s := NewState(nil, ""); s.currentMenu != "anonymous" {
		t.Fatal("a new state without login should init to anonymous")
	}
	// logged_in
	if s := NewState(nil, "test"); s.currentMenu != "logged_in" {
		t.Fatal("a new state with login should init to logged_in")
	}
}
