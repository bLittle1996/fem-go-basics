package io

import (
	"os"
)

func WriteFile(fileName string, contents string) error {
	return os.WriteFile(fileName, []byte(contents), 0644)
}
