package main

import (
	"os"
	"strings"
)

// genAuto will assume that every string that is divisible with 4 is a base64 encoded string and will try to decrypt it.
// it the encryption fails it will not return error but will assume that the string was not encrypted by KMS and continue with the next itteration in the loop.
// Returns error.
func (d *kmsDecrypter) genAuto() string {
	// Create a result channel and count var
	result := make(chan result)
	count := 0

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

		count++
		// Send the key and value + result channel to the decrypt function in a separate go-routine.
		go d.decrypt(&key, &value, result)
	}

	// Loop through and wait for all go-routines to finish their decryption phase.
	for i := 0; i < count; i++ {
		res := <-result

		// Here is where auto gets dangerous. If decryption fails we will assume that it's a none KMS encrypted value!
		if res.err != nil {
			continue
		}

		format := strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(*d.output, "{KEY}", res.key, -1), "{VAL}", res.val, -1), "{CRLF}", "\r\n", -1), "{LF}", "\n", -1), "{TAB}", "\t", -1)
		ret = ret + format

	}

	return ret
}
