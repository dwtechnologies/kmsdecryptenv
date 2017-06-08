package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	d, err := config()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result := map[string]string{}
	// Determine if we should automatically decrypt all envs or just those with our specified marker.
	switch {
	case *d.auto:
		result = d.decrypter.DecryptEnvAuto()
	default:
		result, err = d.decrypter.DecryptEnv(*d.marker)
		// If we get any errors, print the error and exit with code 1.
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Create the output according to our output format.
	ret := ""
	for key, val := range result {
		format := strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(*d.output, "{KEY}", key, -1), "{VAL}", val, -1), "{CRLF}", "\r\n", -1), "{LF}", "\n", -1), "{TAB}", "\t", -1)
		ret = ret + format
	}

	// Print the result
	fmt.Printf(ret)
}
