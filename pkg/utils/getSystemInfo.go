package utils

import (
	"os"
	"os/user"
)

func GetSystemInfo() (string, string, string, string, error) {
	// Get hostname
	hostname, err := os.Hostname()
	if err != nil {
		return "", "", "", "", err
	}

	// Get username
	usr, err := user.Current()
	if err != nil {
		return hostname, "", "", "", err
	}

	// Get working directory
	workingDir, err := os.Getwd()
	if err != nil {
		return hostname, usr.Username, "", "", err
	}

	// Get user home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return hostname, usr.Username, workingDir, "", err
	}

	return hostname, usr.Username, workingDir, homeDir, nil
}
