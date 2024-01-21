package executer

import (
	"fmt"
	"os/exec"

	"github.com/MhamzaAhmad/commence/constants"
	"github.com/MhamzaAhmad/commence/utils"
)

func Execute(name string) {
	if checkForConfig(name) {
		cmd := exec.Command("/bin/bash", fmt.Sprintf("%s/%s.sh", constants.STORAGE_DIR, name))

		err := cmd.Run()

		if err != nil {
			panic(err)
		}
	}
}

func checkForConfig(name string) bool {
	dir := constants.STORAGE_DIR

	utils.CheckForDirectory(dir)

	if exist := utils.CheckForFile(fmt.Sprintf("%s/%s.sh", dir, name)); !exist {
		return false
	}

	return true
}
