package utils

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func ExecInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate the command and the argument
	args := strings.Split(input, " ")

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		// 'cd' to home dir with empty path not yet supported.
		if len(args) < 2 {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return errors.New("error getting user home directory")
			}
			return os.Chdir(homeDir)
		}

		// Change the directory and return the error.
		return os.Chdir(args[1])
	case "exit":
		// Exit the shell
		os.Exit(0)
	}

	// Prepare the command to execute
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
