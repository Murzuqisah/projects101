package main_test

import (
	"fmt"
	"numbers/prime"
	"testing"
)

// testValues := int(7, 15, 34, 2, 3)
// expected := string("true", "false", "false", "true", "true")

type primeTest struct {
	n        int
	expected bool
}

var primeTests = []primeTest{
	primeTest{7, true},
	primeTest{15, false},
	primeTest{34, false},
	primeTest{2, true},
	primeTest{3, true},
}

func TestMain(t *testing.T) {
	for _, tc := range primeTests {
		t.Run(fmt.Sprintf("primeTest_%d", tc.n), func(t *testing.T) {
			result := prime.IsPrime(tc.n)
			if result != tc.expected {
				t.Errorf("For primeTest %d, expected a string output, but got %t", tc.n, tc.expected)
			}
		})
	}

}

// func TestMain(t *testing.T) {

// 	//IsPrime := prime.IsPrime(testValues)

// 	// range through the test values(PrimeTests)
// 	for _, test := range primeTests {
// 		output := prime.IsPrime(test.n)
// 		expected := test.expected
// 		if output != expected {
// 			t.Fatalf("Expected %s not equal to output %s", expected, output)
// 		}
// 	}

// }
