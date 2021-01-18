package client

import (
	"bufio"
	"os"

	"github.com/pkg/errors"
)

// getValidInput returns the selected menu command as a string or an error
func (s *State) getValidInput() (string, error) {
	menu := s.config.Menus[s.currentMenu]

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadByte()
		if err != nil {
			return "", errors.Wrapf(err, "Could not read byte from stdin")
		}
		for _, menuEntry := range menu.MenuEntries {
			if []byte(menuEntry.Key)[0] == input {
				return menuEntry.Action, nil
			}
		}
	}
}
