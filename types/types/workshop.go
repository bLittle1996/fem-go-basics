package types

import (
	"fmt"
	"time"
)

type Workshop struct {
	Course
	Date time.Time
}

func NewWorkshop(date time.Time, course Course) Workshop {
	var workshop Workshop

	workshop.Date = date
	workshop.Course = course

	return workshop

}

func (workshop Workshop) String() string {
	return fmt.Sprintf(
		"New workshop from %s! Checkout %s on %s %d %d at %d:%02d",
		workshop.Instructor.Name,
		workshop.Name,
		workshop.Date.Month(),
		workshop.Date.Day(),
		workshop.Date.Year(),
		workshop.Date.Hour(),
		workshop.Date.Minute(),
	)
}
