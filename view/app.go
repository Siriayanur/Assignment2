package view

import (
	"fmt"
	"os"

	"github.com/Siriayanur/Assignment2/controller"
	"github.com/Siriayanur/Assignment2/exceptions"
)

type App struct {
	data controller.Data
}

func RunApp() {
	var choice int
	data, err := controller.CreateStudentArray()
	app := App{data: *data}
	exceptions.CreateErrorStatements()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		displayMenu()
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			err := app.AddStudentDetails()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("User added successfully..")
			}
		case 2:
			err := app.DisplayStudentDetails()
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			err := app.DeleteStudentDetails()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("User deleted successfully..")
			}
		case 4:
			err := app.SaveStudentDetails()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("All changes saved..")
			}
		case 5:
			err := app.ConfirmExit()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Terminating the Program - Success")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
			os.Exit(1)
		}
	}
}

func (app *App) AddStudentDetails() error {
	fullName, age, rollNumber, address := readUserDetails()
	coursesEnrolled := readCourseDetails()
	err := app.data.AddStudent(fullName, age, rollNumber, address, coursesEnrolled)
	return err
}
func (app *App) DisplayStudentDetails() error {
	sortParameter, sortOrder := getSortParameter()
	err := app.data.DisplayStudents(sortParameter, sortOrder)
	if err != nil {
		return err
	}
	for i := 0; i < len(app.data.Students); i++ {
		app.data.Students[i].SingleStudentDetail()
	}
	return nil
}
func (app *App) DeleteStudentDetails() error {
	var target string
	fmt.Println("Enter the roll number whose record to be deleted :: ")
	fmt.Scanln(&target)
	err := app.data.DeleteStudent(target)
	if err != nil {
		return err
	}
	// remove the entry from map.
	delete(app.data.TrackRollNum, target)
	return nil
}
func (app *App) SaveStudentDetails() error {
	return app.data.SaveStudent()
}
func (app *App) ConfirmExit() error {
	// Ask if they want to save data
	var choice string
	fmt.Println("Do you want to save the changes ? y/n")
	fmt.Scanln(&choice)
	if choice == "y" || choice == "yes" {
		err := app.data.SaveStudent()
		if err != nil {
			return err
		}
	}
	return nil
}
