package parser

import (
	"fmt"

	commandParser "github.com/MhamzaAhmad/commence/parser/command"
	"github.com/MhamzaAhmad/commence/parser/generator"
	"github.com/MhamzaAhmad/commence/reader"
	"github.com/MhamzaAhmad/commence/utils"
	"github.com/MhamzaAhmad/commence/writer"
)

func Parse(path string) {
	sequence := reader.Read(path)

	name := sequence.Name
	cmds := sequence.Commands

	utils.DeleteIfExists(fmt.Sprintf("%s.sh", name))

	for _, cmd := range cmds {
		parsed, err := commandParser.Parse(fmt.Sprintf("%s.yaml", name), []byte(cmd))

		if err != nil {
			panic(err)
		}

		parsedCommand := parsed.(*commandParser.Command)

		generated := generator.Generate(*parsedCommand)

		writer.Write(fmt.Sprintf("%s.sh", name), generated)

		writer.Write(fmt.Sprintf("%s.sh", name), "\n")
	}
}
