package types

import "fmt"

type Instructor struct {
	Id    int
	Name  string
	score int
}

func NewInstructor(id int, name string) Instructor {
	return Instructor{Id: id, Name: name}
}

func (instructor *Instructor) SetScore(score int) {
	if instructor == nil {
		return
	}

	instructor.score = score
}

func (instructor Instructor) String() string {
	return fmt.Sprintf("%s (%d)", instructor.Name, instructor.score)
}
