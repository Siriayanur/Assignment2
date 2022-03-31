package exceptions

import (
	"errors"
	"fmt"
)

var ErrMap = make(map[string]string)

var ErrInvalidFileOperation = errors.New("file Error ")
var ErrInvalidStudentDetails = errors.New("student Details Error ")
var ErrInvalidInput = errors.New("invalid Input ")
var ErrInvalidSortParameter = errors.New("invalid Sort Parameter ")

func CreateErrorStatements() {
	ErrMap["createFile"] = " cannot create file with the file name "
	ErrMap["removeFile"] = " cannot remove file -- check permissions "
	ErrMap["writeFile"] = " cannot write to the file -- check permissions "
	ErrMap["readFile"] = " cannot read the file "
	ErrMap["unmarshal"] = " cannot unmarshal the file contents "
	ErrMap["rollNumExists"] = " roll number already exists "
	ErrMap["rollNumNotExists"] = " cannot delete record - roll number not exists "
	ErrMap["int"] = " enter only integer data "
	ErrMap["sortParameter"] = "Allowed :: 1.Full Name 2.Roll Number 3.Age 3.Address, 1.Asc 2.Desc only acceptable "
}

func InvalidOperation(statement string, errKind error) error {
	return fmt.Errorf("%w :: %s", errKind, ErrMap[statement])
}
