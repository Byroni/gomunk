package main

import (
	"os"

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
