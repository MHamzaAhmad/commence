package parser

import (
	"fmt"

	"github.com/MhamzaAhmad/commence/reader"
)

func Parse(path string) {
	sequence := reader.Read(path)

	cmds := sequence.Commands

	fmt.Println(cmds)
}
