package lastword

import (
	"os"

	"github.com/01-edu/z01"
)

func LastWd() {
	str := os.Args[1]
	for str[len(str)-1] == ' ' {
		str = str[:len(str)-1]
	}

	res := ""
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' ' {
			res = str[i+1:]
			break
		}
	}
	if res == "" {
		res = str
	}

	for _, ch := range res {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
