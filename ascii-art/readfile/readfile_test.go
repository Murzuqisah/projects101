package readfile

import "testing"

func TestReadFile(t *testing.T) {
	content, err := ReadFile("standard.txt")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}
	expectedContent := "Content of the file"
	if content != expectedContent {
		t.Errorf("Expected content %s, but got %s", expectedContent, content)
	}
}
