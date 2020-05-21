package gomunk

import (
	"fmt"
	"os"
	"path/filepath"
)

type goMunk struct {
	fileStoreHandler FileStoreHandler
}

func GoMunk(provider string) *goMunk {
	switch provider {
	case "aws":
		return &goMunk{
			fileStoreHandler: AWSFileStore(),
		}
	default:
		panic(fmt.Errorf("%s", "Must specify valid file store provider"))
	}
}

func (gm *goMunk) ListFiles() {
	gm.fileStoreHandler.ListFiles()
}

func (gm *goMunk) UploadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	key := filepath.Base(path)

	fmt.Printf("Uploading file: %s...\n", key)

	err = gm.fileStoreHandler.UploadFile(key, file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Uploaded file to file store")
}
