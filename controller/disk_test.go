package controller

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckValidFileHappyFlow(t *testing.T) {
	fileName := "main_operations.go"
	isValid := checkValidFile(fileName)
	require.Equal(t, isValid, true)
}
func TestCheckValidFileBadFlow(t *testing.T) {
	fileName := "random.json"
	isValid := checkValidFile(fileName)
	require.Equal(t, isValid, false)
}
func TestCreateFileHappyFlow(t *testing.T) {
	filePointer, _ := createFile()
	require.NotNil(t, filePointer)
}
