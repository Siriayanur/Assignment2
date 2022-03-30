package view

import (
	"fmt"
	"os"

	"github.com/Siriayanur/Assignment2/exceptions"
	"github.com/Siriayanur/Assignment2/model"
)

func RunApp() {
	var choice int
	data, err := model.CreateStudentArray()
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
			err := data.AddStudentDetails()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			break
		case 2:
			err := data.DisplayStudents()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			break
		case 3:
			err := data.DeleteStudentDetails()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			break
		case 4:
			err := data.SaveStudentDetails()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			break
		case 5:
			err := data.ConfirmExit()
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
