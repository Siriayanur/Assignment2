package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStudentDetailsValidationHappyFlow(t *testing.T) {
	var courseEnrolled []Course
	courseEnrolled = append(courseEnrolled, Course{Name: "A", Code: "1"})
	courseEnrolled = append(courseEnrolled, Course{Name: "B", Code: "2"})
	courseEnrolled = append(courseEnrolled, Course{Name: "C", Code: "3"})
	courseEnrolled = append(courseEnrolled, Course{Name: "D", Code: "4"})
	testStudents := []struct {
		fullName       string
		age            int
		rollNumber     string
		address        string
		courseEnrolled []Course
	}{
		{fullName: "siri", age: 20, rollNumber: "1si18cs008", address: "C40 Sampige Apartments", courseEnrolled: courseEnrolled},
		{fullName: "siri890", age: 19, rollNumber: "1si18cs139", address: "C404 Sevanthi Apartments", courseEnrolled: courseEnrolled},
		{fullName: "anonymous", age: 100, rollNumber: "1si18cs134", address: "B41 Prestige Blore", courseEnrolled: courseEnrolled},
		{fullName: "raama", age: 100, rollNumber: "1s1000", address: "ayodhya1000", courseEnrolled: courseEnrolled},
	}
	for _, student := range testStudents {
		st := Student{FullName: student.fullName, RollNumber: student.rollNumber, Age: student.age, Address: student.address, CourseEnrolled: student.courseEnrolled}
		invalidStudent := st.ValidateStudentDetails()
		require.Nil(t, invalidStudent)
	}
}

func TestStudentDetailsValidationSadFlow(t *testing.T) {
	var courseEnrolled []Course
	courseEnrolled = append(courseEnrolled, Course{Name: "A", Code: "1"})
	courseEnrolled = append(courseEnrolled, Course{Name: "B", Code: "2"})
	courseEnrolled = append(courseEnrolled, Course{Name: "C", Code: "3"})
	courseEnrolled = append(courseEnrolled, Course{Name: "D", Code: "4"})
	testStudents := []struct {
		fullName       string
		age            int
		rollNumber     string
		address        string
		courseEnrolled []Course
	}{
		{fullName: "si", age: 20, rollNumber: "1si18cs008", address: "C40 Sampige Apartments", courseEnrolled: courseEnrolled},
		{fullName: "a", age: 19, rollNumber: "1si18cs139", address: "C404 Sevanthi Apartments", courseEnrolled: courseEnrolled},
		{fullName: "kerr", age: -100, rollNumber: "1si18cs139", address: "B41 Prestige Blore", courseEnrolled: courseEnrolled},
		{fullName: "zooo", age: -9, rollNumber: "1s1000", address: "ayodhya1000", courseEnrolled: courseEnrolled},
	}
	for _, student := range testStudents {
		st := Student{FullName: student.fullName, RollNumber: student.rollNumber, Age: student.age, Address: student.address, CourseEnrolled: student.courseEnrolled}
		invalidStudent := st.ValidateStudentDetails()
		require.NotNil(t, invalidStudent)
	}
}
