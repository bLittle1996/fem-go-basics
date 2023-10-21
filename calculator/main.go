package main

import "fmt"

func main() {
	var operation string
	var numberA, numberB int

	fmt.Println("Advanced Calculator")
	fmt.Println("===================")
	fmt.Println("\nChoose operation:")
	fmt.Scanf("%s", &operation)
	fmt.Println("Choose first num:")
	fmt.Scanf("%d", &numberA)
	fmt.Println("Choose second num:")
	fmt.Scanf("%d", &numberB)

	switch operation {
	case "add":
		fmt.Printf("%d + %d = %d\n", numberA, numberB, numberA+numberB)
	case "sub":
		fallthrough
	case "subtract":
		fmt.Printf("%d - %d = %d\n", numberA, numberB, numberA-numberB)
	case "multiply":
		fmt.Printf("%d * %d = %d\n", numberA, numberB, numberA*numberB)
	case "divide":
		fmt.Printf("%d / %d = %d\n", numberA, numberB, numberA/numberB)
	default:
		panic("invalid operation")
	}
}
