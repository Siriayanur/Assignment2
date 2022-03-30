package model

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Student struct {
	FullName       string
	Age            int
	RollNumber     string
	Address        string
	CourseEnrolled []Course
}

func (s Student) SingleStudentDetail() {
	fmt.Println("| Student Details |")
	fmt.Printf("Full Name : %s | Age : %d | Roll Number : %s \n", s.FullName, s.Age, s.RollNumber)
	fmt.Println("| Courses enrolled |")
	for i := 0; i < len(s.CourseEnrolled); i++ {
		fmt.Printf("Course Name :: %s, Course Code :: %s\n", s.CourseEnrolled[i].name, s.CourseEnrolled[i].code)
	}
}
func (s Student) ValidateStudentDetails() error {

	return validation.ValidateStruct(&s,
		validation.Field(&s.FullName, validation.Required, validation.Length(4, 30)),
		validation.Field(&s.Age, validation.Required, validation.Min(1)),
		validation.Field(&s.RollNumber, validation.Required, validation.Length(1, 50)),
		validation.Field(&s.Address, validation.Required, validation.Length(1, 50)),
		validation.Field(&s.CourseEnrolled, validation.Required),
	)
}
