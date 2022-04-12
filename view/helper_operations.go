package view

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Siriayanur/Assignment2/model"
	"github.com/Siriayanur/Assignment2/utils"
)

func displayMenu() {
	fmt.Println("| MENU |")
	fmt.Println("1. ADD USER DETAILS")
	fmt.Println("2. DISPLAY USER DETAILS |")
	fmt.Println("3. DELETE USER DETAILS |")
	fmt.Println("4. SAVE USER DETAILS |")
	fmt.Println("5. EXIT")
}

// read input from user.
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
func readCourseDetails() []model.Course {
	// Display all the available courses
	model.DisplayCourses()

	var coursesEnrolled []model.Course
	isEnrolled := []int{0, 0, 0, 0, 0, 0}
	count := 0
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
		coursesEnrolled = append(coursesEnrolled, model.CourseDetails[courseIndex])
		// Break if sufficient courses are enrolled
		if count >= utils.NumberOfSubjectsPerStudent {
			break
		}
	}
	return coursesEnrolled
}
func getSortParameter() (int, int) {
	var choice1, choice2 int
	fmt.Println("Enter the parameter to consider for sorting : ")
	fmt.Println("1.Full Name 2.Roll Number 3.Age 4.Address : ")
	fmt.Scan(&choice1)
	fmt.Println("Enter the order for sorting 1-asce 2-desc : ")
	fmt.Scan(&choice2)
	return choice1, choice2
}
