package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
)

type kmsDecrypter struct {
	auto    *bool
	session *session.Session
	region  *string
	marker  *string
}

func main() {
	d := config()

	// Determine if we should automatically decrypt all envs or just those with our specified marker.
	switch {
	case *d.auto:
		fmt.Printf(d.genAuto())
	default:
		ret, err := d.genFromMarked()
		// If we get any errors, print the error and exit with code 1.
		if err != nil {
			panic(err)
		}

		fmt.Printf(ret)
	}
}
