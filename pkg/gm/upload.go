package gm

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/byroni/gomunk/pkg/aws"
)

// NewGoMunk creates a new GoMunkUpload client
func GoMunkUpload(path string) *goMunkUpload {
	return &goMunkUpload{
		path:    path,
		keyName: filepath.Base(path),
	}
}

// GoMunkUpload type struct
type goMunkUpload struct {
	path    string
	keyName string
}

// UploadFile uploads a file
func (gm *goMunkUpload) UploadFile() {
	file, err := gm.OpenFile(gm.path)
	if err != nil {
		log.Fatal("Unable to locate file", err)
	}
	defer file.Close()

	fmt.Printf("Uploading file: %s...\n", gm.path)

	client := aws.GoMunkS3()

	_, err = client.UploadFile(gm.path, file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Uploaded file to file store")
}

// OpenFile is a utility function that opens a file for reading
func (gm *goMunkUpload) OpenFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return file, err
	}

	return file, nil
}
