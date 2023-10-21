package main

import (
	"fmt"
	"front-end-masters/go-basics/in-out/data"
)

func init() {
	fmt.Println("Hello from main.go")
}

func main() {
	data.SomeDict["haha"] = data.FunnyNumber == 42069
	fmt.Printf("%v\n", data.SomeDict)
	printSomethingIdk()
}
