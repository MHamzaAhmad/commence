package generator

import "fmt"

func Open(process string, loc string) string {
	switch process {
		case "vscode":
			return fmt.Sprintf("VSCODE_CWD='%s' open -n -b 'com.microsoft.VSCode' --args . ;", loc)
		default:
			panic("Invalid process")
	}
}
