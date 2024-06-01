package splitlines

import (
	"fmt"
	"testing"
)

func TestSplitLines(t *testing.T) {
	content := "line1\nline2\nline3\n"
	lines := SplitLines(content, "standard.txt")
	expectedLines := []string{"line1", "line2", "line3", ""}
	if !compareslices(lines, expectedLines) {
		fmt.Println("Unable to split lines")
	}
}

func compareslices(str1, str2 []string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for i := range str1 {
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}