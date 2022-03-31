package controller

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Siriayanur/Assignment2/exceptions"
	"github.com/Siriayanur/Assignment2/model"
	"github.com/Siriayanur/Assignment2/utils"
)

func checkValidFile(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func createFile() (*os.File, error) {
	// If exists, remove
	if checkValidFile(utils.FileName) {
		err := os.Remove(utils.FileName)
		if err != nil {
			return nil, exceptions.InvalidOperation("removeFile", exceptions.ErrInvalidFileOperation)
		}
	}
	// create new file with same name
	filePointer, err := os.Create(utils.FileName)
	if err != nil {
		return nil, err
	}
	return filePointer, nil
}
func ReadDataFromDisk() ([]model.Student, error) {
	if !checkValidFile(utils.FileName) {
		// create new file to store the data
		_, err := createFile()
		if err != nil {
			return nil, err
		}
	}
	openFile, err := os.Open(utils.FileName)
	if err != nil {
		return nil, err
	}
	defer openFile.Close()
	studentDataRaw, err := ioutil.ReadAll(openFile)
	if err != nil {
		return nil, err
	}
	if len(studentDataRaw) == 0 {
		var emptyData []model.Student
		return emptyData, nil
	}

	var studentDataMarshal []model.Student
	// fmt.Println("data from file : ", studentDataRaw)
	err = json.Unmarshal(studentDataRaw, &studentDataMarshal)
	if err != nil {
		return nil, err
	}
	return studentDataMarshal, nil
}
func SaveDataToDisk(students []model.Student) error {
	// convert to json
	marshalData, err := json.Marshal(students)
	if err != nil {
		return err
	}
	// overwrite the data --> delete old file,create new file with the current data
	filePointer, err := createFile()
	if err != nil {
		return err
	}
	defer filePointer.Close()
	_, err = filePointer.Write(marshalData)
	if err != nil {
		return err
	}
	return nil
}