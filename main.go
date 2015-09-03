package main

import (
	"flag"
	"fmt"
	"github.com/masterlink720/shellcheckgo/api"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var (
		filename string
		err      error
	)

	// Setup arg flags
	flag.StringVar(&filename, "script", "", "Path to the script to be verified")
	flag.Parse()

	// Fail
	if filename == "" {
		ShowHelp()
	}

	// Resolve the path
	filename, err = GetPath(filename)

	// Try fetching the script
	script, err := GetScript(filename)
	if err != nil {
		log.Fatal("Fatal error while reading the script", err)
		os.Exit(1)
	}

	// Success! now let's do something with it
	api.Check(script)

	//fmt.Print(fmt.Sprintf("Script successfully read into memory!:\n\n%s", script))

}

/**
 * Exit the program and print a quick help message
 */
func ShowHelp() {
	flag.Usage()
	fmt.Print("\n")

	os.Exit(1)
}

/**
 * Attempt to verify / normalize the script path provided
 */
func GetPath(filename string) (string, error) {
	return filepath.Abs(filepath.Clean(filename))
}

/**
 * Fetch the contents of the script
 */
func GetScript(path string) (string, error) {
	f, err := os.Open(path)

	if err != nil {
		return "", err
	}

	// Defer closing the file
	defer f.Close()

	// We'll use a byte buffer to read the file
	var result []byte
	buffer := make([]byte, 100)
	for {
		n, err := f.Read(buffer[0:])
		result = append(result, buffer[0:n]...)

		// oh noes!
		if err != nil {
			// no worries, just end of file
			if err == io.EOF {

				// All done!
				break
			}

			// Fail
			return "", err
		}
	}

	// Success!
	return string(result), nil
}
