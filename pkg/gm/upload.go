package gm

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/byroni/gomunk/pkg/aws"
)

// NewGoMunk creates a new GoMunkUpload client
func NewGoMunkUpload(path string) *GoMunkUpload {
	return &GoMunkUpload{
		path:    path,
		keyName: filepath.Base(path),
	}
}

// GoMunkUpload type struct
type GoMunkUpload struct {
	path    string
	keyName string
}

// UploadFile uploads a file
func (gm *GoMunkUpload) UploadFile() {
	bucket := "gomunk-file-store"

	file, err := gm.OpenFile(gm.path)
	if err != nil {
		log.Fatal("Unable to locate file", err)
	}
	defer file.Close()

	fmt.Println("Uploading file:", gm.path)
	uploader := aws.NewUploader(aws.NewSession())

	options := aws.UploadInput(&bucket, &gm.path, file)

	_, err = uploader.Upload(options)
	if err != nil {
		panic(err)
	}

	fmt.Println("Uploaded file to file store")
}

// OpenFile is a utility function that opens a file for reading
func (gm *GoMunkUpload) OpenFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return file, err
	}

	return file, nil
}
