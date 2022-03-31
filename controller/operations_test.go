package controller

import (
	"testing"

	"github.com/Siriayanur/Assignment2/model"
)

func GenerateTestData() *Data {
	data := Data{}
	var courseEnrolled []model.Course
	courseEnrolled = append(courseEnrolled, model.Course{Name: "A", Code: "1"})
	courseEnrolled = append(courseEnrolled, model.Course{Name: "B", Code: "2"})
	courseEnrolled = append(courseEnrolled, model.Course{Name: "C", Code: "3"})
	courseEnrolled = append(courseEnrolled, model.Course{Name: "D", Code: "4"})
	testStudents := []struct {
		fullName       string
		age            int
		rollNumber     string
		address        string
		courseEnrolled []model.Course
	}{
		{fullName: "siri", age: 20, rollNumber: "12", address: "tumkur1", courseEnrolled: courseEnrolled},
		{fullName: "siri890", age: 19, rollNumber: "123", address: "ashoka nagar tumkur", courseEnrolled: courseEnrolled},
		{fullName: "siri123", age: 100, rollNumber: "125", address: "tumkur2", courseEnrolled: courseEnrolled},
		{fullName: "raama", age: 100, rollNumber: "s100", address: "tumkur0", courseEnrolled: courseEnrolled},
	}
	for _, student := range testStudents {
		st := model.Student{FullName: student.fullName, RollNumber: student.rollNumber, Age: student.age, Address: student.address, CourseEnrolled: student.courseEnrolled}
		data.Students = append(data.Students, st)
	}
	data.populateMap(data.Students)
	return &data
}
func TestDisplayStudentHappyFlow(t *testing.T) {
	data := GenerateTestData()
	// sort by full name, asc
	err := data.DisplayStudentsHelper(1, 1)
	if !StudentsOrderByFullName(data.Students) || err != nil {
		t.Errorf("Expected order of records by full name Ascending order, got error : %v", err)
	}
	// sort by age, asc
	err0 := data.DisplayStudentsHelper(3, 1)
	if !StudentsOrderByAge(data.Students) || err0 != nil {
		t.Errorf("Expected order of records by age Ascending order, got error : %v", err0)
	}
	// sort by address, asc
	err1 := data.DisplayStudentsHelper(4, 1)
	if !StudentsOrderByAddress(data.Students) || err1 != nil {
		t.Errorf("Expected order of records by address Ascending order, got error : %v", err1)
	}
	// sort by rollNumber, asc
	err2 := data.DisplayStudentsHelper(2, 1)
	if !StudentsOrderByRollNumber(data.Students) || err2 != nil {
		t.Errorf("Expected order of records by Roll Number Ascending order, got error : %v", err2)
	}
	// sort by rollNumber, desc
	err3 := data.DisplayStudentsHelper(2, 2)
	if !StudentsOrderByRollNumberDesc(data.Students) || err3 != nil {
		t.Errorf("Expected order of records by Roll Number Descending order got error : %v", err3)
	}
}
func StudentsOrderByFullName(students []model.Student) bool {
	for i := 0; i < len(students)-1; i++ {
		if students[i].FullName > students[i+1].FullName {
			return false
		}
	}
	return true
}
func StudentsOrderByAge(students []model.Student) bool {
	for i := 0; i < len(students)-1; i++ {
		if students[i].Age > students[i+1].Age {
			return false
		}
	}
	return true
}
func StudentsOrderByRollNumber(students []model.Student) bool {
	for i := 0; i < len(students)-1; i++ {
		if students[i].RollNumber > students[i+1].RollNumber {
			return false
		}
	}
	return true
}
func StudentsOrderByRollNumberDesc(students []model.Student) bool {
	for i := 0; i < len(students)-1; i++ {
		if students[i].RollNumber < students[i+1].RollNumber {
			return false
		}
	}
	return true
}
func StudentsOrderByAddress(students []model.Student) bool {
	for i := 0; i < len(students)-1; i++ {
		if students[i].Address > students[i+1].Address {
			return false
		}
	}
	return true
}
func TestDeleteStudentDetailsHappyFlow(t *testing.T) {
	data := GenerateTestData()
	target := "123"
	err := data.DeleteStudentDetailsHelper(target)
	if err != nil {
		t.Errorf("Some error in deleting student record : %s", target)
	}
}
func TestDeleteStudentDetailsBadFlow(t *testing.T) {
	data := GenerateTestData()
	target := "1"
	err := data.DeleteStudentDetailsHelper(target)
	if err == nil {
		t.Errorf("Expected to not delete the student record : %s, got error : %v", target, err)
	}
}
func TestSaveAndRetrieve(t *testing.T) {
	data := GenerateTestData()
	err := data.SaveStudentDetails()
	if err != nil {
		t.Errorf("Could not write to disk, got error %v", err)
	}
	readData, err := ReadDataFromDisk()
	if err != nil {
		t.Errorf("Could not read data from disk, got error %v", err)
	}
	if len(readData) != len(data.Students) {
		t.Errorf("Mismatch in data records saved and retrieved : expected %d records, got %d records !", len(data.Students), len(readData))
	}
	for i := 0; i < len(readData); i++ {
		if readData[i].RollNumber != data.Students[i].RollNumber ||
			readData[i].Age != data.Students[i].Age ||
			readData[i].Address != data.Students[i].Address ||
			readData[i].FullName != data.Students[i].FullName {
			t.Errorf("Mismatch in data records values : expected %v records, got %v records !", data.Students[i], readData[i])
		}
	}
}
