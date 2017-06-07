package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
)

type kmsDecrypter struct {
	auto    *bool
	session *session.Session
	region  *string
	marker  *string
	output  *string
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
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf(ret)
	}
}
