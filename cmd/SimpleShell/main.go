package main

import (
	"SimpleShell/pkg/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		hostname, username, workingDir, homeDir, err := utils.GetSystemInfo()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if workingDir == homeDir {
			workingDir = "~"
		} else {
			workingDir = utils.GetDirBaseName(workingDir)
		}

		fmt.Print("[", username, "@", hostname, " ", workingDir, "]$ ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = utils.ExecInput(input)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
