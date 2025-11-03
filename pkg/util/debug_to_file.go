package util

import (
	"os"
)

func AppendToFile(text string) error {
	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile("/tmp/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Append data to the file
	if _, err := file.WriteString(text + "\n"); err != nil {
		return err
	}

	return nil
}
