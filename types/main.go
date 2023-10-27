package main

import (
	"fmt"
	"front-end-masters-go-basics/types/types"
	"strings"
	"time"
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

	instructor := types.NewInstructor(4, "Gregolomo")
	instructor.SetScore(68)

	fmt.Println(instructor)

	course := types.Course{Id: 404, Name: "How to write 404s", Instructor: instructor}

	workshop := types.NewWorkshop(time.Now().Add(time.Hour*24), types.Course{Id: 420, Name: "Embed Deez", Instructor: instructor})

	var courses [2]types.Printable

	courses[0] = course
	courses[1] = workshop

	for _, course := range courses {
		fmt.Println(course)
	}
}
