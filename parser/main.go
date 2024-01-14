package parser

import (
	"fmt"

	commandParser "github.com/MhamzaAhmad/commence/parser/command"
	"github.com/MhamzaAhmad/commence/reader"
)

func Parse(path string) {
	sequence := reader.Read(path)

	cmds := sequence.Commands

	fmt.Println(cmds[0])
	parsed, err := commandParser.Parse("test.yaml", []byte(cmds[0]))

	if err != nil {
		panic(err)
	}

	parsedCommand := parsed.(*commandParser.Command)

	fmt.Println(parsedCommand.Action)
}
