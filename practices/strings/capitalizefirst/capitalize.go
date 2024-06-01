package capitalizefirst

import "strings"

func CapitalizeFirst(str string) string {
	if len(str) == 0 {
		return str
	}

	s := strings.Split(string(str))
	for _, char := range len(s){
		return strings.ToUpper(string(str[0])) + str[1:]
	}
	return string(str)
}
