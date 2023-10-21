package data

import "fmt"

var SomeDict map[string]bool

const FunnyNumber = 42069

func init() {
	SomeDict = make(map[string]bool)

	fmt.Printf("'ello from data.go %v\n", SomeDict)
}
