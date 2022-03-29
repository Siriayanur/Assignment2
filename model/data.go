package model

import (
	"bufio"
	"fmt"
	"os"
	"reflect"

	"github.com/Siriayanur/Assignment2/utils"
)

type Data struct {
	Students []Student
}

func CreateStudentArray() (*Data, error) {
	data := Data{}
	var err error

	// var data []Student
	data.Students, err = ReadDataFromDisk()
	fmt.Print(reflect.ValueOf(data).Kind())
	if err != nil {
		fmt.Println("no error")
	}
	return &data, nil
}

func (d *Data) AddStudentDetails() {
	fullName, age, rollNumber, address := readUserDetails()
	var coursesEnrolled []Course = readCourseDetails()
	student := Student{FullName: fullName, Address: address, Age: age, RollNumber: rollNumber, CourseEnrolled: coursesEnrolled}
	invalidStudent := student.ValidateStudentDetails()
	if invalidStudent != nil {
		fmt.Println(invalidStudent.Error())
		os.Exit(1)
	}
	d.Students = append(d.Students, student)
}

func (d *Data) DisplayStudents() {
	fmt.Println("Came here")
	for i := 0; i < len(d.Students); i++ {
		d.Students[i].SingleStudentDetail()
	}
	fmt.Println("Came here")
}

func (d *Data) DeleteStudentDetails() {

}
func (d *Data) SaveStudentDetails() {
	err := SaveDataToDisk(d.Students)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func (d *Data) ConfirmExit() {
}
func readUserDetails() (string, int, string, string) {
	var fullName string
	var address string
	var age int
	var rollNumber string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter fullname : ")
	scanner.Scan()
	fullName = scanner.Text()
	fmt.Println("Enter age : ")
	fmt.Scanf("%d", &age)
	fmt.Println("Enter roll number : ")
	scanner.Scan()
	rollNumber = scanner.Text()
	fmt.Println("Enter address : ")
	scanner.Scan()
	address = scanner.Text()
	return fullName, age, rollNumber, address
}

func readCourseDetails() []Course {
	// Display all the available courses
	DisplayCourses()

	var coursesEnrolled []Course
	isEnrolled := []int{0, 0, 0, 0, 0, 0}
	var count = 0
	for {
		fmt.Println("Enter course number :: ")
		var courseIndex int
		_, err := fmt.Scanf("%d", &courseIndex)
		// Check if user entered int data
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Check if the entered course number is in correct range
		if courseIndex < 0 || courseIndex >= utils.TotalNumberOfSubjects {
			fmt.Println("Not a valid course Number")
			continue
		}
		// Check if course is already enrolled
		if isEnrolled[courseIndex] != 0 {
			fmt.Println("This course is already enrolled!!")
			continue
		}
		isEnrolled[courseIndex] = 1
		count++
		coursesEnrolled = append(coursesEnrolled, CourseDetails[courseIndex])
		// Break if sufficient courses are enrolled
		if count >= utils.NumberOfSubjectsPerStudent {
			break
		}
	}
	return coursesEnrolled
}
