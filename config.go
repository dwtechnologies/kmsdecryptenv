package main

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	autoEnv   = "KMS_AUTO_DECRYPT"
	markerEnv = "KMS_MARKER"
	regionEnv = "KMS_AWS_REGION"
)

func config() *kmsDecrypter {
	// Fetch AWS_REGION from ENV
	r := os.Getenv(regionEnv)
	if r == "" {
		r = "eu-west-1" // Default to eu-west-1 if no region is specified.
	}

	// Get KMS Marker
	m := os.Getenv(markerEnv)
	if m == "" {
		m = "KMS_MARKER"
	}

	// Get AUTO value
	a := false
	if strings.ToLower(os.Getenv(autoEnv)) == "true" {
		a = true
	}

	// Use the default credential provider, it will check in the following order (1. ENV VARS, 2. CONFIG FILE, 3. EC2 ROLE).
	// In most cases we will use the EC2 role for providing access to the KMS key.
	c := credentials.NewCredentials(nil)
	cfg := &aws.Config{
		Credentials: c,
		Region:      &r,
	}

	return &kmsDecrypter{
		session: session.Must(session.NewSession(cfg)),
		region:  &r,
		marker:  &m,
		auto:    &a,
	}
}
