package main

import (
	"fmt"
	"path/filepath"

	"github.com/baoist/resource/command"
	"github.com/baoist/resource/generator"
)

func main() {
	cmd := command.GetCommand()

	run(cmd)
}

func run(command command.Command) {
	switch command.Name {
	case "generate":
		config, err := filepath.Abs("./config.yml")
		if err != nil {
			panic(err)
		}

		generator.Generate(config)
	default:
		fmt.Println("The command", fmt.Sprintf("'%v'", command.Name), "is not supported.")
	}
}
