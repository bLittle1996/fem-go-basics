package main

import (
	"fmt"
	"front-end-masters/go-basics/in-out/data"
)

func init() {
	fmt.Println("Hello from main.go")
}

func doSomeMaths(number float64) (float64, float64) {
	return number * 4.20, number * 0.69
}

func main() {
	data.SomeDict["haha"] = data.FunnyNumber == 42069
	fmt.Printf("%v\n", data.SomeDict)
	printSomethingIdk()
	a, b := doSomeMaths(40.0)
	fmt.Printf("%f, %f\n", a, b)
}
