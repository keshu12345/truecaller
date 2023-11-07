package util

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadPrefixesFromFile(t *testing.T) {
	// Create a temporary directory and write sample files for testing
	tempDir := t.TempDir()
	file1Path := filepath.Join(tempDir, "file1.txt")
	file2Path := filepath.Join(tempDir, "file2.txt")

	// Create and write contents to sample files
	file1Content := "prefix1\nprefix2\nprefix3\n"
	file2Content := "prefix4\nprefix5\n"

	if err := os.WriteFile(file1Path, []byte(file1Content), 0644); err != nil {
		t.Fatalf("Failed to create file1: %v", err)
	}

	if err := os.WriteFile(file2Path, []byte(file2Content), 0644); err != nil {
		t.Fatalf("Failed to create file2: %v", err)
	}

	// Call the function to read prefixes
	prefixes := ReadPrefixesFromFile(tempDir)

	// Check if the expected prefixes are found
	expectedPrefixes := []string{"prefix1", "prefix2", "prefix3", "prefix4", "prefix5"}
	for i, prefix := range expectedPrefixes {
		if prefixes[i] != prefix {
			t.Errorf("Expected prefix '%s', but got '%s'", prefix, prefixes[i])
		}
	}
}

func TestMatcherPrefixesList(t *testing.T) {
	// Create a temporary directory and write a sample prefixes file for testing
	tempDir := t.TempDir()
	prefixesFilePath := filepath.Join(tempDir, samplePrefixesFile)
	prefixesContent := "sample_prefix1\nsample_prefix2\nsample_prefix3\n"

	if err := os.WriteFile(prefixesFilePath, []byte(prefixesContent), 0644); err != nil {
		t.Fatalf("Failed to create sample prefixes file: %v", err)
	}

	// Call the function to read the sample prefixes file
	prefixes, err := MatcherPrefixesList(prefixesFilePath)

	// Check if there was an error
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check if the expected prefixes are found
	expectedPrefixes := []string{"sample_prefix1", "sample_prefix2", "sample_prefix3"}
	for i, prefix := range expectedPrefixes {
		if prefixes[i] != prefix {
			t.Errorf("Expected prefix '%s', but got '%s'", prefix, prefixes[i])
		}
	}
}
