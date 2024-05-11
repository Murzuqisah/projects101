package validate

import "strings"

func IsValid(word string) bool {
	word = strings.ToLower(word)
	vowel := string("a,e,i,o,u,A,E,I,O,U")
	num := string("0,1,2,3,4,5,6,7,8,9")
	consonants := string("bcdfghjklmnpqrstvwxyz")
	badChar := string("~!@#$%^&*_+")

	if len(word) < 3 {
		return false
	}

	if len(word) >= 3 {
		if strings.ContainsAny(word, num) {
			if strings.ContainsAny(word, badChar) {
				return false
			}
			if !strings.ContainsAny(word, vowel) {
				return false
			}
			if !strings.ContainsAny(word, consonants) {
				return false
			}
		}
		if strings.ContainsAny(word, badChar) {
			return false
		}
		if !strings.ContainsAny(word, num) && !strings.ContainsAny(word, vowel) {
			return false
		}
	}
	return true
}
