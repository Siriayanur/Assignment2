package controller

import (
	"fmt"
	"sort"

	"github.com/Siriayanur/Assignment2/exceptions"
	"github.com/Siriayanur/Assignment2/model"
)

type Data struct {
	Students     []model.Student
	TrackRollNum map[string]bool
}

// main operations.
func (d *Data) AddStudentDetails() error {
	fullName, age, rollNumber, address := readUserDetails()
	coursesEnrolled := readCourseDetails()
	student := model.Student{FullName: fullName, Age: age, RollNumber: rollNumber, Address: address, CourseEnrolled: coursesEnrolled}
	ErrInvalidStudent := student.ValidateStudentDetails()
	if ErrInvalidStudent != nil {
		return ErrInvalidStudent
	}
	if d.TrackRollNum[student.RollNumber] {
		return exceptions.InvalidOperation("rollNumExists", exceptions.ErrInvalidStudentDetails)
	}
	d.TrackRollNum[student.RollNumber] = true
	d.Students = append(d.Students, student)
	return nil
}
func (d *Data) DisplayStudentsHelper(sortParameter int, sortOrder int) error {
	if sortOrder != 1 && sortOrder != 2 {
		return exceptions.InvalidOperation("sortParameter", exceptions.ErrInvalidSortParameter)
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
	if sortOrder == 2 {
		d.Students = reverseArray(d.Students)
	}
	return nil
}
func (d *Data) DisplayStudents() error {
	// ask for sorting parameter and order
	sortParameter, sortOrder := getSortParameter()
	err := d.DisplayStudentsHelper(sortParameter, sortOrder)
	if err != nil {
		return err
	}
	for i := 0; i < len(d.Students); i++ {
		d.Students[i].SingleStudentDetail()
	}
	return nil
}
func (d *Data) DeleteStudentDetailsHelper(target string) error {
	// check existence of roll num
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
		d.Students = []model.Student{}
	case targetIndex == len(d.Students)-1:
		d.Students = d.Students[:targetIndex]
	default:
		d.Students = append(d.Students[:targetIndex], d.Students[targetIndex+1:]...)
	}
	return nil
}
func (d *Data) DeleteStudentDetails() error {
	var target string
	fmt.Println("Enter the roll number whose record to be deleted :: ")
	fmt.Scanln(&target)
	err := d.DeleteStudentDetailsHelper(target)
	if err != nil {
		return err
	}
	// remove the entry from map.
	delete(d.TrackRollNum, target)
	return nil
}
func (d *Data) SaveStudentDetails() error {
	// sort
	sort.Slice(d.Students, func(i int, j int) bool {
		// if fullName same, then sort with rollNum
		if d.Students[i].FullName == d.Students[j].FullName {
			return d.Students[i].RollNumber < d.Students[j].RollNumber
		}
		return d.Students[i].FullName < d.Students[j].FullName
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
