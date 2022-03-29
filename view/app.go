package view

import (
	"fmt"
	"os"

	"github.com/Siriayanur/Assignment2/model"
)

func RunApp() {
	var choice int
	data, err := model.CreateStudentArray()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		displayMenu()
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			data.AddStudentDetails()
			break
		case 2:
			data.DisplayStudents()
			break
		case 3:
			data.DeleteStudentDetails()
			break
		case 4:
			data.SaveStudentDetails()
			break
		case 5:
			data.ConfirmExit()
		default:
			fmt.Println("Invalid choice")
			os.Exit(1)
		}
	}
}
