package utils

import (
	"os"
	"strings"
)

// ReadInput reads the entire input file and returns it as a string
func ReadInput(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadLines reads the input file and returns it as a slice of lines
func ReadLines(filepath string) ([]string, error) {
	data, err := ReadInput(filepath)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(data), "\n"), nil
}
