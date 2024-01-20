package executer

import (
	commandParser "github.com/MhamzaAhmad/commence/parser/command"
	executer "github.com/MhamzaAhmad/commence/parser/executer/open"
)

func Execute(
	command commandParser.Command,
) {
	switch command.Action {
		case "open":
			executer.Open(command.Process.(string), command.Location.(string));
		default:
			panic("Invalid action")
	}
}
