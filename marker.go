package main

import (
	"fmt"
	"os"
	"strings"
)

// genFromMarked will decrypt all envs that contain has the specified marker in it's name.
// Returns error.
func (d *kmsDecrypter) genFromMarked() (string, error) {
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

		unecrypted, err := d.kmsDecrypt(&value)
		if err != nil {
			return ret, err
		}

		ret = fmt.Sprintf("%v%v=%v\n", ret, newkey, *unecrypted)
	}

	return ret, nil
}
