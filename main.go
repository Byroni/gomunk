package main

func main() {
	path := "./testFile.txt"

	client := NewGoMunk(path)

	client.UploadFile()
}
