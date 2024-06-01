package checkarguments_test

import (
	"testing"

	"ascii/checkarguments"
)

type argsTest struct {
	args    string
	wantErr bool
}

var argTests = []argsTest{
	{"hello€", false},
	{"invalid\x00Char", true},
}

func TestCheckArguments(t *testing.T) {
	for _, tt := range argTests {
		err := checkarguments.CheckArguments(tt.args)
		if (err != nil) != tt.wantErr {
			t.Errorf("CheckArguments(%q) error = %v, wantErr %v", tt.args, err, tt.wantErr)
		}
	}
}
