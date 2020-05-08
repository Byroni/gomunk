package gm

import "github.com/byroni/gomunk/pkg/aws"

func Get(key string) {
	client := aws.GoMunkS3()
	client.Get(key)
}
