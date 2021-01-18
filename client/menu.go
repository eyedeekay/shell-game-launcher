package client

import "fmt"

func (s *State) displayMenu() {
	menu := s.config.Menus[s.currentMenu]
	fmt.Print("\033[2J") // clear the screen
	fmt.Printf("%s\n\n", menu.Banner)
	if s.login == "" {
		fmt.Print("Not logged in.\n\n")
	} else {
		fmt.Printf("Logged in as: %s\n\n", s.login)
	}
	for i := 0; i < len(menu.MenuEntries); i++ {
		fmt.Printf("%s) %s\n", menu.MenuEntries[i].Key, menu.MenuEntries[i].Label)
	}
}
