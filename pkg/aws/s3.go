package aws

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type GoMunkS3Handler interface {
	Bucket() string
	UploadFile(string, *os.File) (*s3manager.UploadOutput, error)
	Get(string)
	List()
}

func GoMunkS3() GoMunkS3Handler {
	client := &goMunkS3{
		bucket: "gomunk-file-store",
	}

	client.session = NewSession()
	client.s3Session = s3.New(client.session)
	client.uploader = s3manager.NewUploader(client.session)

	return client
}

type goMunkS3 struct {
	bucket    string
	session   *session.Session
	s3Session *s3.S3
	uploader  *s3manager.Uploader
}

func (gm *goMunkS3) Bucket() string {
	return gm.bucket
}

func (gm *goMunkS3) UploadFile(key string, body *os.File) (*s3manager.UploadOutput, error) {
	input := &s3manager.UploadInput{
		Bucket: aws.String(gm.bucket),
		Key:    aws.String(key),
		Body:   body,
	}

	output, err := gm.uploader.Upload(input)
	if err != nil {
		return output, err
	}
	return output, nil
}

func (gm *goMunkS3) List() {
	input := &s3.ListObjectsInput{
		Bucket:  aws.String(gm.bucket),
		MaxKeys: aws.Int64(100),
	}

	result, err := gm.s3Session.ListObjects(input)
	if err != nil {
		PrintAWSError(err)
	}

	const MiB float64 = 1048576
	var totalSize int64
	for _, v := range result.Contents {
		totalSize += *v.Size
		fmt.Printf("%s | %-20s | %sB\n", v.LastModified, *v.Key, strconv.FormatInt(*v.Size, 10))
	}
	fmt.Printf("\nTotal bucket size: %sMiB\n", strconv.FormatFloat(float64(totalSize)/MiB, 'g', 1, 64))
}

func (gm *goMunkS3) Get(key string) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(gm.bucket),
		Key:    aws.String(key),
	}
	result, err := gm.s3Session.GetObject(input)
	if err != nil {
		PrintAWSError(err)
	}

	file, err := ioutil.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filepath.Base(key), file, 0644)
	if err != nil {
		panic(err)
	}
}
