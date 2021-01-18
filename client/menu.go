package client

import "fmt"

func (s *State) displayMenu() {
	menu := s.config.Menus[s.currentMenu]
	fmt.Print("\033[2J") // clear the screen
	fmt.Println(menu.Banner)
	output := ""
	prefix := ""
	for i := 0; i < menu.XOffset; i++ {
		prefix += " "
	}
	for i := 0; i < menu.YOffset; i++ {
		output += "\n"
	}
	fmt.Printf("%s", output)
	for i := 0; i < len(menu.MenuEntries); i++ {
		fmt.Printf("%s%s) %s\n", prefix, menu.MenuEntries[i].Key, menu.MenuEntries[i].Label)
	}
}
