package gomunk

import "os"

type FileStoreHandler interface {
	UploadFile(string, *os.File) error
	GetFile(string) error
	ListFiles()
}
