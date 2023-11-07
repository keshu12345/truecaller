package util

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	logger "github.com/sirupsen/logrus"
)

const (
	samplePrefixesFile = "sample_prefixes.txt"
)

func ReadPrefixesFromFile(directory string) []string {
	var prefixes []string

	err := filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				logger.Errorf("Error opening file %s: %v\n", path, err)
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				prefix := strings.TrimSpace(scanner.Text())
				if prefix != "" {
					prefixes = append(prefixes, prefix)
				}
			}

			if err := scanner.Err(); err != nil {
				logger.Errorf("Error reading file %s: %v\n", path, err)
			}
		}

		return nil
	})

	if err != nil {
		logger.Errorf("Error walking directory: %v\n", err)
	}

	return prefixes
}

func MatcherPrefixesList(fileName string) ([]string, error) {
	// Define a command-line flag for the file path
	// filePath := flag.String("file", "", "Path to the target file")
	// flag.Parse()

	if fileName == "" {
		fileName = samplePrefixesFile
	}

	p, err := filepath.Abs(fileName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return []string{}, fmt.Errorf("Error: %v\n", err)
	}

	prefixes := ReadPrefixesFromFile(p)
	return prefixes, nil

}
