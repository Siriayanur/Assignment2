package model

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/Siriayanur/Assignment2/utils"
)

type Data struct {
	Students     []Student
	TrackRollNum map[string]bool
}

func CreateStudentArray() (*Data, error) {
	data := Data{}
	var err error
	data.Students, err = ReadDataFromDisk()
	// populate TrackRollNum map
	data.populateMap(data.Students)
	// fmt.Print(reflect.ValueOf(data).Kind())
	if err != nil {
		fmt.Println("could not read data from disk ")
		return nil, err
	}
	return &data, nil
}
func (d *Data) populateMap(students []Student) {
	for i := 0; i < len(students); i++ {
		d.TrackRollNum[students[i].RollNumber] = true
	}
}
func (d *Data) AddStudentDetails() error {
	fullName, age, rollNumber, address := readUserDetails()
	var coursesEnrolled []Course = readCourseDetails()
	student := Student{FullName: fullName, Address: address, Age: age, RollNumber: rollNumber, CourseEnrolled: coursesEnrolled}
	invalidStudent := student.ValidateStudentDetails()
	if invalidStudent != nil {
		return invalidStudent
	}
	if d.TrackRollNum[student.RollNumber] {
		return errors.New("Roll Number already exists")
	}
	d.Students = append(d.Students, student)
	return nil
}
func getSortParameter() (int, int) {
	var choice1, choice2 int
	fmt.Println("Enter the parameter to consider for sorting : ")
	fmt.Println("1.Full Name 2.Roll Number 3.Age 4.Address : ")
	fmt.Scanln(&choice1)
	fmt.Println("Enter the order for sorting 1-asce 2-desc : ")
	fmt.Scanln(&choice2)
	return choice1, choice2
}
func (d *Data) DisplayStudents() {
	// ask for sorting parameter and order
	sortParameter, sortOrder := getSortParameter()
	switch sortParameter {
	case 1:
		sort.Slice(d.Students, func(i int, j int) bool {
			if d.Students[i].FullName == d.Students[j].FullName {
				return d.Students[i].RollNumber < d.Students[j].RollNumber
			}
			return d.Students[i].FullName < d.Students[j].FullName
		})
	case 2:
		sort.Slice(d.Students, func(i int, j int) bool {
			return d.Students[i].RollNumber < d.Students[j].RollNumber
		})
	case 3:
		sort.Slice(d.Students, func(i int, j int) bool {
			return d.Students[i].Age < d.Students[j].Age
		})
	case 4:
		sort.Slice(d.Students, func(i int, j int) bool {
			if d.Students[i].Address == d.Students[j].Address {
				return d.Students[i].RollNumber < d.Students[j].RollNumber
			}
			return d.Students[i].Address < d.Students[j].Address
		})
	default:
		fmt.Println("Invalid Sort Choice")
		os.Exit(1)
	}
	if sortOrder != 1 && sortOrder != 2 {
		fmt.Println("Invalid Ordering Choice")
		os.Exit(1)
	}
	if sortOrder == 2 {
		d.Students = reverseArray(d.Students)
	}
	for i := 0; i < len(d.Students); i++ {
		d.Students[i].SingleStudentDetail()
	}
}
func reverseArray(s []Student) []Student {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
func (d *Data) DeleteStudentDetails() {
	// to be done
}
func (d *Data) SaveStudentDetails() {
	// sort
	sort.Slice(d.Students, func(i int, j int) bool {
		// if fullName same, then sort with rollNum
		if d.Students[i].FullName == d.Students[i].FullName {
			return d.Students[i].RollNumber < d.Students[i].RollNumber
		}
		return d.Students[i].FullName < d.Students[i].FullName
	})
	// save the sorted data to disk
	err := SaveDataToDisk(d.Students)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func (d *Data) ConfirmExit() {
	// to be done
	// Ask if they want to save data
	d.SaveStudentDetails()
	// else exit
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
