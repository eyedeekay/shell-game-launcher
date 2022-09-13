package client

import (
	"github.com/adyxax/shell-game-launcher/config"
)

type State struct {
	config      *config.Config
	currentMenu string
	login       string
}

func NewState(config *config.Config, login string) *State {
	cs := State{
		config:      config,
		currentMenu: "anonymous",
		login:       login,
	}
	if login != "" {
		cs.currentMenu = "logged_in"
	}
	return &cs
}
