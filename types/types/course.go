package types

import "fmt"

type Course struct {
	Id         int
	Name       string
	Instructor Instructor
	Duration   float64
}

func (course Course) String() string {
	return fmt.Sprintf("Course: %s by %s", course.Name, getOr(course.Instructor.Name, "some teacher idk"))
}

func (course Course) SignUpOrSomethingIdk() error {
	return nil
}

func getOr(str string, fallback string) string {
	if str == "" {
		return fallback
	}

	return str
}
