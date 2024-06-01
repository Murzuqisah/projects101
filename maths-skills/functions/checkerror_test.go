package functions

import "testing"

func TestCheckNilError(t *testing.T) {
	type args struct {
		err       error
		customMsg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckNilError(tt.args.err, tt.args.customMsg); (err != nil) != tt.wantErr {
				t.Errorf("CheckNilError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
