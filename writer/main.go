package writer

import (
	"fmt"
	"os"

	"github.com/MhamzaAhmad/commence/constants"
	"github.com/MhamzaAhmad/commence/utils"
)


func Write(filename string, command string) {
	dir := constants.STORAGE_DIR

	utils.CheckForDirectory(dir)

	f, err := os.OpenFile(fmt.Sprintf("%s/%s", dir, filename) , os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if file, err := f.Stat(); err == nil {
		if file.Size() == 0 {
			if _, err := f.WriteString("#!/bin/bash\n"); err != nil {
				panic(err)
			}
		}
	}

	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err := f.WriteString(command); err != nil {
		panic(err)
	}
}

