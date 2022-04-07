package controller

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/Siriayanur/Assignment2/exceptions"
	"github.com/Siriayanur/Assignment2/model"
	"github.com/Siriayanur/Assignment2/utils"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}
func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
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
	ciphertext, err := ioutil.ReadAll(openFile)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) == 0 {
		var emptyData []model.Student
		return emptyData, nil
	}
	// decrypt the data before displaying
	plaintext := decrypt(ciphertext, secrete_code)
	var studentDataMarshal []model.Student
	// fmt.Println("data from file : ", studentDataRaw)
	err = json.Unmarshal(plaintext, &studentDataMarshal)
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
	// encrypt the marshalData before storing
	ciphertext := encrypt([]byte(marshalData), secrete_code)
	// overwrite the data --> delete old file,create new file with the current data
	filePointer, err := createFile()
	if err != nil {
		return err
	}
	defer filePointer.Close()
	_, err = filePointer.Write(ciphertext)
	if err != nil {
		return err
	}
	return nil
}
