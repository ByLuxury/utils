package util

import (
	"encoding/base64"
	"os"
	"io/ioutil"
	"log"
)

// DecodeBase64String
func DecodeBase64String(src string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(src)
	return data, err
}

// EncodeBase64String ...
func EncodeBase64String(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	buf := make([]byte, 5000000)
	n, err := f.Read(buf)
	if err != nil {
		return "", err
	}
	src := base64.StdEncoding.EncodeToString(buf[:n])
	return src, nil
}

// Base64ToImage Read a base64 txt file and decode to imgage
func Base64ToImage(filename string) error {
	// read file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("read file failed %v", err)
		return err
	}
	// decode
	dist, err := base64.StdEncoding.DecodeString(string(file))
	//fmt.Println(string(file))
	if err != nil {
		log.Fatalf("decode base64 string failed:%v", err)
		return err
	}
	// write a new file
	f, err := os.OpenFile(filename+".png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalf("open a file instance failed:%v", err)
		return err
	}
	defer f.Close()
	_, err = f.Write(dist)
	return err
}
