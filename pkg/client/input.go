package client

import (
	"bufio"
	"fmt"
	"os"
)

// getValidInput returns the selected menu command as a string or an error
func (s *State) getValidInput() (string, error) {
	menu := s.config.Menus[s.currentMenu]

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadByte()
		if err != nil {
			return "", fmt.Errorf("Could not read byte from stdin : %w", err)
		}
		for _, menuEntry := range menu.MenuEntries {
			if []byte(menuEntry.Key)[0] == input {
				return menuEntry.Action, nil
			}
		}
	}
}
