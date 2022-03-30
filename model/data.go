package model

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/Siriayanur/Assignment2/exceptions"
	"github.com/Siriayanur/Assignment2/utils"
)

type Data struct {
	Students     []Student
	TrackRollNum map[string]bool
}

// main operations
func (d *Data) AddStudentDetails() error {
	fullName, age, rollNumber, address := readUserDetails()
	var coursesEnrolled []Course = readCourseDetails()
	student := Student{FullName: fullName, Address: address, Age: age, RollNumber: rollNumber, CourseEnrolled: coursesEnrolled}
	invalidStudent := student.ValidateStudentDetails()
	if invalidStudent != nil {
		return invalidStudent
	}
	if d.TrackRollNum[student.RollNumber] {
		return exceptions.InvalidOperation("rollNumExists", exceptions.ErrInvalidStudentDetails)
	} else {
		d.TrackRollNum[student.RollNumber] = true
	}
	d.Students = append(d.Students, student)
	return nil
}
func (d *Data) DisplayStudents() error {
	// ask for sorting parameter and order
	sortParameter, sortOrder, err := getSortParameter()
	if err != nil {
		return err
	}
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
		return exceptions.InvalidOperation("sortParameter", exceptions.ErrInvalidSortParameter)
	}
	if sortOrder != 1 && sortOrder != 2 {
		return exceptions.InvalidOperation("sortParameter", exceptions.ErrInvalidSortParameter)
	}
	if sortOrder == 2 {
		d.Students = reverseArray(d.Students)
	}
	for i := 0; i < len(d.Students); i++ {
		d.Students[i].SingleStudentDetail()
	}
	return nil
}
func (d *Data) DeleteStudentDetails() error {
	var target string
	fmt.Println("Enter the roll number whose record to be deleted :: ")
	fmt.Scanln(&target)
	// check existance of roll num
	if !d.TrackRollNum[target] {
		return exceptions.InvalidOperation("rollNumNotExists", exceptions.ErrInvalidStudentDetails)
	}
	// find target index
	targetIndex := -1
	for i := 0; i < len(d.Students); i++ {
		if d.Students[i].RollNumber == target {
			targetIndex = i
		}
	}
	switch {
	case len(d.Students) == 1:
		d.Students = []Student{}
	case targetIndex == len(d.Students)-1:
		d.Students = d.Students[:targetIndex]
	default:
		d.Students = append(d.Students[:targetIndex], d.Students[targetIndex+1:]...)
	}
	//remove the entry from map
	delete(d.TrackRollNum, target)
	return nil
}
func (d *Data) SaveStudentDetails() error {
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
		return exceptions.InvalidOperation("writeFile", err)
	}
	return nil
}
func (d *Data) ConfirmExit() error {
	// Ask if they want to save data
	var choice string
	fmt.Println("Do you want to save the changes ? y/n")
	fmt.Scanln(&choice)
	if choice == "y" || choice == "yes" {
		err := d.SaveStudentDetails()
		if err != nil {
			return err
		}
	}
	return nil
}

// helper functions
func getSortParameter() (int, int, error) {
	var choice1, choice2 int
	fmt.Println("Enter the parameter to consider for sorting : ")
	fmt.Println("1.Full Name 2.Roll Number 3.Age 4.Address : ")
	fmt.Scan(&choice1)
	fmt.Println("Enter the order for sorting 1-asce 2-desc : ")
	fmt.Scan(&choice2)
	return choice1, choice2, nil
}
func (d *Data) populateMap(students []Student) {
	d.TrackRollNum = map[string]bool{}
	for i := 0; i < len(students); i++ {
		d.TrackRollNum[students[i].RollNumber] = true
	}
}
func PrintMap(d *Data) {
	fmt.Println("Map data")
	for key, value := range d.TrackRollNum {
		fmt.Printf("%s,%t \n", key, value)
	}
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
func reverseArray(s []Student) []Student {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// read input from user
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
