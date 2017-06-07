package main

import (
	"os"
	"strings"
)

// genAuto will assume that every string that is divisible with 4 is a base64 encoded string and will try to decrypt it.
// it the encryption fails it will not return error but will assume that the string was not encrypted by KMS and continue with the next itteration in the loop.
// Returns error.
func (d *kmsDecrypter) genAuto() string {
	envs := os.Environ()
	ret := ""

	// Range over the envs.
	for _, env := range envs {
		slice := strings.SplitN(env, "=", 2)
		if len(slice) != 2 {
			continue
		}

		key := slice[0]
		value := slice[1]

		// If key isn't divisible with 4 continue
		if (len(key) % 4) != 0 {
			continue
		}

		// Here is where auto gets dangerous. If decryption fails we will assume that it's a none KMS encrypted value!
		unecrypted, err := d.kmsDecrypt(&value)
		if err != nil {
			continue
		}

		format := strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(*d.output, "{KEY}", key, -1), "{VAL}", *unecrypted, -1), "{CRLF}", "\r\n", -1), "{LF}", "\n", -1), "{TAB}", "\t", -1)
		ret = ret + format
	}

	return ret
}
