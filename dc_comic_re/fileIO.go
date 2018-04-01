package main

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/petershen0307/goLogger/logClient"
)

// CreateDir to create folder
func CreateDir(pathName string) {
	if _, err := os.Stat(pathName); os.IsNotExist(err) {
		logClient.Log(logClient.LevelInfo, "Create root dir: %v", pathName)
		if err := os.Mkdir(pathName, os.ModeDir); err != nil {
			panic(err)
		}
	}
}

// IsFileExist to check file exist or not
func IsFileExist(dir, fileName string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		CreateDir(dir)
	}
	// check path exist or not
	// if path exist os.Stat() return err will be nil, can't use os.IsExist()
	/*
		os.IsExist(), os.IsNotExist() reference: https://gist.github.com/mastef/05f46d3ab2f5ed6a6787
		os.IsExist(err) is good for cases when you expect the file to not exist yet,
		but the file actually exists :
	*/
	filePath := filepath.Join(dir, fileName)
	if _, err := os.Stat(filePath); err == nil {
		logClient.Log(logClient.LevelInfo, "File exist: %v", filePath)
		return true
	}
	return false
}

// WriteFile raw date to file
func WriteFile(buf *bytes.Buffer, dir, fileName string) error {
	if IsFileExist(dir, fileName) {
		return nil
	}

	path := filepath.Join(dir, fileName)
	fo, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer fo.Close()
	_, err = buf.WriteTo(fo)
	return err
}
