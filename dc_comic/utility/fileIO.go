package utility

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// CreateDir to create folder
func CreateDir(pathName string) {
	if _, err := os.Stat(pathName); os.IsNotExist(err) {
		log.Println("Create root dir: ", pathName)
		if err := os.Mkdir(pathName, os.ModeDir); err != nil {
			panic(err)
		}
	}
}

// IsFileExist to check file exist or not
func IsFileExist(dir, fileName string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Println("Create dir: ", dir)
		if err := os.Mkdir(dir, os.ModeDir); err != nil {
			panic(err)
		}
	}
	// check path exist or not
	// if path exist os.Stat() return err will be nil, can't use os.IsExist()
	/*
		os.IsExist(), os.IsNotExist() reference: https://gist.github.com/mastef/05f46d3ab2f5ed6a6787
		os.IsExist(err) is good for cases when you expect the file to not exist yet,
		but the file actually exists :
	*/
	path := fmt.Sprintf("%s/%s", dir, fileName)
	if _, err := os.Stat(path); err == nil {
		log.Println("File exist: ", path)
		return true
	}
	return false
}

// WriteFile raw date to file
func WriteFile(reader io.Reader, dir, fileName string) {
	if IsFileExist(dir, fileName) {
		return
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	log.Println("2 Data Length: ", len(data))
	path := fmt.Sprintf("%s/%s", dir, fileName)
	fo, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	fo.Write(data)
}
