package main

import (
	"encoding/base64"

	"github.com/aws/aws-sdk-go/service/kms"
)

type result struct {
	key string
	val string
	err error
}

// decrypt will call on kmsDecrypt using seperate go-channels to make it concurrent.
// will send a result struct to the res channel when completed.
func (d *kmsDecrypter) decrypt(key *string, val *string, res chan<- result) {
	str, err := d.kmsDecrypt(*val)
	result := result{
		key: *key,
		val: str,
	}
	if err != nil {
		result.err = err
	}

	// Send the result to the result channel.
	res <- result
}

// kmsDecrypt will decrypt the encrypted value using KMS and return the unecrypted value in plaintext.
// Returns *string and error.
func (d *kmsDecrypter) kmsDecrypt(str string) (string, error) {
	svc := kms.New(d.session)

	// We want to have the encrypted value in base64. So decode it to to []byte.
	decode, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	// Decrypt using KMS.
	params := &kms.DecryptInput{CiphertextBlob: decode}

	resp, err := svc.Decrypt(params)
	if err != nil {
		return "", err
	}

	return string(resp.Plaintext), nil
}
