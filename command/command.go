package command

import (
	"os"
)

type Command struct {
	Name    string
	Options []string
}

func GetCommand() (command Command) {
	command.Name = os.Args[1]

	for _, option := range os.Args[2:] {
		command.Options = append(command.Options, option)
	}

	return command
}
