package main

import (
	"fmt"
	"os"
	"strings"
)

// genFromMarked will decrypt all envs that contain has the specified marker in it's name.
// It will create a new env var without the marker in it's name and set the original env to empty (for security purposes).
// Returns error.
func (d *kmsDecrypter) genFromMarked() error {
	envs := os.Environ()

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
			return err
		}

		// Set the new, unecrypted value.
		err = os.Setenv(newkey, *unecrypted)
		if err != nil {
			return err
		}

		// Remove the encrypted key for security purposes.
		err = os.Setenv(key, "")
		if err != nil {
			return err
		}
	}

	return nil
}
