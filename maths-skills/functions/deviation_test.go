package functions

import "testing"

func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := StandardDeviation()
			if got != tt.want {
				t.Errorf("StandardDeviation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("StandardDeviation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
