package filesystem

import (
	"os"
)

func HasPath(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true // Path exists
	}
	if os.IsNotExist(err) {
		return false // Path does not exist
	}
	return false // Some other error, like permission issue
}
