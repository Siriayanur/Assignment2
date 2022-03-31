package controller

import "testing"

func TestCheckValidFileHappyFlow(t *testing.T) {
	fileName := "operations.go"
	isValid := checkValidFile(fileName)
	if !isValid {
		t.Errorf("File doesn't exist,expected %t, got %t", true, isValid)
	}
}
func TestCheckValidFileBadFlow(t *testing.T) {
	fileName := "random.json"
	isValid := checkValidFile(fileName)
	if isValid {
		t.Errorf("File exists,expected %t, got %t", true, isValid)
	}
}
func TestCreateFileHappyFlow(t *testing.T) {
	filePointer, _ := createFile()
	if filePointer == nil {
		t.Errorf("Could not create file, expected valid new file pointer, got nil")
	}
}
