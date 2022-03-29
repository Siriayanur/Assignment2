package exceptions

import (
	"errors"
	"fmt"
)

var ErrMap = make(map[string]string)

var errInvalidFileOperation = errors.New("File Error ")

func CreateErrorStatements() {
	ErrMap["createFile"] = " cannot create file with the file name "
	ErrMap["removeFile"] = " cannot remove file -- check permissions "
	ErrMap["readFile"] = " cannot read the file "
	ErrMap["unmarshal"] = " cannot unmarshal the file contents "
}

func InvalidItemParameter(statement string) error {
	return fmt.Errorf("%w :: %s", errInvalidFileOperation, statement)
}
