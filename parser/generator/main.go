package generator

import (
	"fmt"

	commandParser "github.com/MhamzaAhmad/commence/parser/command"
	generator "github.com/MhamzaAhmad/commence/parser/generator/open"
)

func Generate(
	command commandParser.Command,
) string {
	fmt.Println(command)
	switch command.Action {
		case "open":
			return generator.Open(command.Process.(string), command.Location.(string));
		default:
			panic("Invalid action")
	}
}
