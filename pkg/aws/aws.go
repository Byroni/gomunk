package aws

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var bucket = "gomunk-file-store"

func NewSession() *session.Session {
	const region string = "us-east-1"
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		panic(err)
	}

	return session
}

func UploadInput(bucket *string, key *string, body *os.File) *s3manager.UploadInput {
	return &s3manager.UploadInput{
		Bucket: bucket,
		Key:    key,
		Body:   body,
	}
}

func NewUploader(session *session.Session) *s3manager.Uploader {
	return s3manager.NewUploader(session)
}

func ListObjects(session *session.Session) {
	svc := S3Session(session)

	input := &s3.ListObjectsInput{
		Bucket:  &bucket,
		MaxKeys: aws.Int64(100),
	}

	result, err := svc.ListObjects(input)
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

func Get(session *session.Session, key string) {
	svc := S3Session(session)
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	result, err := svc.GetObject(input)
	if err != nil {
		PrintAWSError(err)
	}

	file, err := ioutil.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(filepath.Base(key), file, 0644)
}

func PrintAWSError(err error) {
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
			case s3.ErrCodeNoSuchBucket:
				log.Fatal(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
}

func S3Session(session *session.Session) *s3.S3 {
	return s3.New(session)
}
