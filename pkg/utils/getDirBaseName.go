package utils

import (
	"strings"
)

func GetDirBaseName(path string) string {
	// Check if the path is the root directory
	if path == "/" {
		return "/"
	}

	// Split the path into parts
	parts := strings.Split(path, "/")

	// The last part is the base name of the directory
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return path
}
