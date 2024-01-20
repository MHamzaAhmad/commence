package executer

import (
	"os"
	"os/exec"
)

func Open(process string, loc string) {
	switch process {
	case "vscode":
		loc = os.ExpandEnv(loc)
		cmd := exec.Command("open", "-n", "-b", "com.microsoft.VSCode", "--args", ".")
		cmd.Dir = loc
		cmd.Env = append(os.Environ(), "VSCODE_CWD="+loc)
		err := cmd.Run()

		if err != nil {
			panic(err)
		}
	}
}
