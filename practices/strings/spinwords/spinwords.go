package spinwords

import (
	"strings"
)

func SpinWords(str string) string {
	byte_str := []rune(str)

	for i, char := range byte_str {
		if len(char) >= 5 {
			res := " "
			for _, line := range char {
				res = string(line) + res
			}
			byte_str[i] = res
		}
	}
	return strings.Join(byte_str, " ")
}
