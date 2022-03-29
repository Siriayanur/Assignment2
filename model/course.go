package model

import (
	"fmt"
)

type Course struct {
	name string
	code string
}

var CourseDetails = []Course{
	{name: "A", code: "1"},
	{name: "B", code: "2"},
	{name: "C", code: "3"},
	{name: "D", code: "4"},
	{name: "E", code: "5"},
	{name: "F", code: "6"},
}

func DisplayCourses() {
	for i := 0; i < len(CourseDetails); i++ {
		fmt.Printf("Number :: %d, course :: %s, course code :: %s\n", i, CourseDetails[i].name, CourseDetails[i].code)
	}
}
