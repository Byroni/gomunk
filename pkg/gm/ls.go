package gm

import "github.com/byroni/gomunk/pkg/aws"

func ListObjects() {
	client := aws.GoMunkS3()
	client.List()
}
