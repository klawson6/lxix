package lxix

import "testing"

func TestRomanNumerals(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		in   int
		want string
	}{
		{
			name: "When only thousand digit is present, then return multiple M characters",
			in:   3000,
			want: "MMM",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanNumerals(tt.in); got != tt.want {
				t.Errorf("RomanNumerals() = %v, want %v", got, tt.want)
			}
		})
	}
}
