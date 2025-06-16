package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadFile(path string) ([]string, error) {
	// If successful, methods on the returned file can be used for reading. Method return io.Reader interface
	file, err := os.Open("prices.txt")
	if err != nil {
		return nil, errors.New("Failed to open file.")
	}

	// New scanner reciving input value with type io.Reader interface
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Error handling for scanner
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read line in file.")
	}

	file.Close()
	return lines, nil
}
