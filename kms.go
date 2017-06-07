package main

import (
	"encoding/base64"

	"github.com/aws/aws-sdk-go/service/kms"
)

// kmsDecrypt will decrypt the encrypted value using KMS and return the unecrypted value in plaintext.
// Returns *string and error.
func (d *kmsDecrypter) kmsDecrypt(str *string) (*string, error) {
	svc := kms.New(d.session)

	// We want to have the encrypted value in base64. So decode it to to []byte.
	decode, err := base64.StdEncoding.DecodeString(*str)
	if err != nil {
		return nil, err
	}

	// Decrypt using KMS.
	params := &kms.DecryptInput{CiphertextBlob: decode}

	resp, err := svc.Decrypt(params)
	if err != nil {
		return nil, err
	}

	plainText := string(resp.Plaintext)
	return &plainText, nil
}
