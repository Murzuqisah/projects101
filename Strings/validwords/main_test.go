package main_test

import (
	"fmt"
	"testing"
	"valid/validate"
)

type validTest struct {
	word     string
	expected bool
}

var validTests = []validTest{
	validTest{"234Adas", true},
	validTest{"aya", true},
	validTest{"b3", false},
	validTest{"a3$e", false},
	validTest{"BCD", false},
	validTest{"1234", false},
}

func TestMain(t *testing.T) {
	for _, test := range validTests {
		t.Run(fmt.Sprintf("validWord_%s", test.word), func(t *testing.T) {
			result := validate.IsValid(test.word)
			if result != test.expected {
				t.Errorf("For validWord %s, expected a boolean output, but got %t", test.word, test.expected)
			}
		})
	}
}
