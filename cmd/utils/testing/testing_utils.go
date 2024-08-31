package testing_utils

import (
	"os"
	"testing"
)

func DeleteFile(filePath string, t *testing.T) {
	err := os.Remove(filePath)
	if err != nil {
		t.Errorf("Failed to delete test file: %v", err)
	}
}
