package io

import (
	"os"
)

func ReadFile(fileName string) (string, error) {
	bytes, error := os.ReadFile(fileName)

	if error != nil {
		return "", error
	}

	contents := string(bytes)

	return contents, nil
}
