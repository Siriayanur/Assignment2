package model

import (
	"fmt"
)

type Course struct {
	Name string
	Code string
}

var CourseDetails = []Course{
	{Name: "A", Code: "1"},
	{Name: "B", Code: "2"},
	{Name: "C", Code: "3"},
	{Name: "D", Code: "4"},
	{Name: "E", Code: "5"},
	{Name: "F", Code: "6"},
}

func DisplayCourses() {
	for i := 0; i < len(CourseDetails); i++ {
		fmt.Printf("Number :: %d, course :: %s, course Code :: %s\n", i, CourseDetails[i].Name, CourseDetails[i].Code)
	}
}
