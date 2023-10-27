package main

import (
	"fmt"
	"strings"
)

type email string

func (email email) domain() string {
	parts := strings.Split(email.string(), "@")

	return parts[1]
}

func (email email) string() string {
	return string(email)
}

func main() {
	var email email = "test@fake.ca"

	fmt.Println(email.domain())

}
