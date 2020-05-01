package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// NewGoMunk creates a new GoMunk client
func NewGoMunk(path string) *GoMunk {
	return &GoMunk{
		path:    path,
		keyName: filepath.Base(path),
	}
}

// GoMunk type struct
type GoMunk struct {
	path    string
	keyName string
}

// UploadFile uploads a file
func (gm *GoMunk) UploadFile() {
	fmt.Println("Uploading file:", gm.path)

	file, err := gm.OpenFile(gm.path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	uploader := NewUploader(NewSession())

	options := UploadInput(&bucket, &gm.path, file)

	_, err = uploader.Upload(options)
	if err != nil {
		panic(err)
	}

	fmt.Println("Uploaded file to file store")
}

// OpenFile is a utility function that opens a file for reading
func (gm *GoMunk) OpenFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return file, err
	}

	return file, nil
}
