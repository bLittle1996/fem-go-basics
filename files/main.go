package main

import (
	"fmt"
	"front-end-masters/go-basics/files/internal"
	"os"
)

func main() {
	cwd, _ := os.Getwd()
	contents, error := io.ReadFile(cwd + "/data/data.txt")

	if error != nil {
		fmt.Printf("%v\n", error)
	} else {
		fmt.Printf("%s\n", contents)
	}

	err := io.WriteFile(cwd+"/data/write-me.txt", "I really like writing text to files!")

	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
