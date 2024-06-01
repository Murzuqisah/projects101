package escapesequences

import (
	"fmt"
	"os"
	"strings"
)

// read the string from the arguments
func EscapeSequences(args []string) string {
	arguments := args[1]
	escapeSeq := []string{"\\t", "\\b", "\\a", "\\r", "\\'", "\\v", "\\f", "\\x20"}

	if arguments == "" || arguments == "\n" {
		return ""
	}
	for _, esc := range escapeSeq {
		if args[1] == esc {
			fmt.Println("Escape sequences")
			os.Exit(0)
		}
	}
	return strings.ReplaceAll(arguments, "\\n", "\n")
}
