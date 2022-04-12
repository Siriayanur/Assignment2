package controller

import (
	"fmt"

	"github.com/Siriayanur/Assignment2/controller/disk"
	"github.com/Siriayanur/Assignment2/model"
)

func CreateStudentArray() (*Data, error) {
	data := Data{}
	var err error
	data.Students, err = disk.ReadDataFromDisk()
	// populate TrackRollNum map.
	data.populateMap(data.Students)
	if err != nil {
		fmt.Println("could not read data from disk ")
		return nil, err
	}
	return &data, nil
}
func reverseArray(s []model.Student) []model.Student {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
