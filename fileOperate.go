package utils

import (
	"io/ioutil"
	"os"
	"io"
)

// ReadFile ...
func ReadFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	return data, err
}

// WriteFile ...
func WriteFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0666)
}

// DeleteFile ...
func DeleteFile(filename string) error {
	err := os.Remove(filename)
	return err
}

// OpenFile ...
func OpenFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	return f, err
}

// CopyFile ...
func CopyFile(file *os.File,filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, file)
	return err
}
