package escapesequences

import (
	"fmt"
	"testing"
)

func TestEscapeSequences(t *testing.T) {
	args := []string{"program", "argument1"}
    arguments := EscapeSequences(args)
    expected := "argument1"
    if arguments != expected {
		fmt.Println("Argument not found")
    }
}
