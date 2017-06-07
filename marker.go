package main

import (
	"fmt"
	"os"
	"strings"
)

// genFromMarked will decrypt all envs that contain has the specified marker in it's name.
// Returns error.
func (d *kmsDecrypter) genFromMarked() (string, error) {
	// Create a result channel and count var
	result := make(chan result)
	count := 0

	envs := os.Environ()
	ret := ""

	for _, env := range envs {
		slice := strings.SplitN(env, "=", 2)
		if len(slice) != 2 {
			continue
		}

		newkey := ""
		key := slice[0]
		value := slice[1]

		// If marker wasn't found, continue to next env var.
		if !strings.Contains(key, *d.marker) {
			continue
		}

		count++
		surrounding := fmt.Sprintf("_%v_", *d.marker)
		leading := fmt.Sprintf("_%v", *d.marker)
		trailing := fmt.Sprintf("%v_", *d.marker)

		switch {
		case strings.Contains(key, surrounding):
			newkey = strings.Replace(key, surrounding, "", 1)

		case strings.Contains(key, leading):
			newkey = strings.Replace(key, leading, "", 1)

		case strings.Contains(key, trailing):
			newkey = strings.Replace(key, trailing, "", 1)

		default:
			newkey = strings.Replace(key, *d.marker, "", 1)
		}

		// Send the key and value + result channel to the decrypt function in a separate go-routine.
		go d.decrypt(&newkey, &value, result)

	}

	// Loop through and wait for all go-routines to finish their decryption phase.
	for i := 0; i < count; i++ {
		res := <-result

		if res.err != nil {
			return ret, res.err
		}

		format := strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(*d.output, "{KEY}", res.key, -1), "{VAL}", res.val, -1), "{CRLF}", "\r\n", -1), "{LF}", "\n", -1), "{TAB}", "\t", -1)
		ret = ret + format

	}

	return ret, nil
}
