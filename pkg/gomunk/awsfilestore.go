package gomunk

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type awsFileStore struct {
	bucket string
}

func AWSFileStore() FileStoreHandler {
	// Get bucket from config by default and allow override by flag
	return &awsFileStore{
		bucket: Config.AWS_BUCKET,
	}
}

func (fs *awsFileStore) UploadFile(key string, body *os.File) error {
	input := &s3manager.UploadInput{
		Bucket: aws.String(fs.bucket),
		Key:    aws.String(key),
		Body:   body,
	}

	session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		panic(err)
	}

	uploader := s3manager.NewUploader(session)

	_, err = uploader.Upload(input)
	if err != nil {
		return err
	}

	return nil
}

func (fs *awsFileStore) GetFile(path string) error {
	input := &s3.GetObjectInput{
		Bucket: aws.String(fs.bucket),
		Key:    aws.String(path),
	}

	session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		return err
	}
	s3Session := s3.New(session)
	result, err := s3Session.GetObject(input)
	if err != nil {
		return err
	}

	file, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Base(filepath.Base(path)), file, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (fs *awsFileStore) ListFiles() {
	input := &s3.ListObjectsInput{
		Bucket:  aws.String(fs.bucket),
		MaxKeys: aws.Int64(100),
	}

	session, err := session.NewSession(&aws.Config{
		Region: aws.String(Config.AWS_REGION),
	})
	if err != nil {
		panic(err)
	}

	s3Session := s3.New(session)

	result, err := s3Session.ListObjects(input)
	if err != nil {
		panic(err)
	}

	const MiB float64 = 1048576
	var totalSize int64
	for _, v := range result.Contents {
		totalSize += *v.Size
		fmt.Printf("%s | %-20s | %sB\n", v.LastModified, *v.Key, strconv.FormatInt(*v.Size, 10))
	}
	fmt.Printf("\nTotal bucket size: %sMiB\n", strconv.FormatFloat(float64(totalSize)/MiB, 'g', 1, 64))
}
