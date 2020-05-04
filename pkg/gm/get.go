package gm

import "github.com/byroni/gomunk/pkg/aws"

func Get(key string) {
	session := aws.NewSession()
	aws.Get(session, key)
}
