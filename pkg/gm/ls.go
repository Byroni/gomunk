package gm

import "github.com/byroni/gomunk/pkg/aws"

var bucket = "gomunk-file-store"

func ListObjects() {
	session := aws.NewSession()
	aws.ListObjects(session)
}
