package main

import (
	"os"
	"strings"

	"github.com/dwtechnologies/kmsdecrypt"
)

const (
	autoEnv   = "KMS_AUTO_DECRYPT"
	markerEnv = "KMS_MARKER"
	defMarker = "KMS_DECRYPT"
	regionEnv = "KMS_AWS_REGION"
	outputVal = "KMS_OUTPUT"
	defOutput = "{KEY}={VAL}{LF}"
)

type decrypter struct {
	decrypter *kmsdecrypt.KmsDecrypter
	auto      *bool
	marker    *string
	output    *string
}

func config() (*decrypter, error) {
	// Get AWS Region and create kmsdecrypt
	r := os.Getenv(regionEnv)
	if r == "" {
		r = "eu-west-1" // Default to eu-west-1 if no region is specified.
	}

	decrypt, err := kmsdecrypt.New(r)
	if err != nil {
		return nil, err
	}

	// Get KMS Marker
	m := os.Getenv(markerEnv)
	if m == "" {
		m = defMarker
	}

	// Get AUTO value
	a := false
	if strings.ToLower(os.Getenv(autoEnv)) == "true" {
		a = true
	}

	// Get Output format
	o := os.Getenv(outputVal)
	if o == "" {
		o = defOutput
	}

	return &decrypter{
		decrypter: decrypt,
		auto:      &a,
		marker:    &m,
		output:    &o,
	}, nil
}
