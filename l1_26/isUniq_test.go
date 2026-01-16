package main

import "testing"

func Test_isUniq(t *testing.T) {
	tests := []struct {
		name string
		input string
		want bool
	}{
		{
			name: "all unique",
			input: "abcd",
			want: true,
		},
		{
			name: "all unique UTF-8",
			input: "ФЫвА",
			want: true,
		},
		{
			name: "duplicated",
			input: "abcdA",
			want: false,
		},
		{
			name: "duplicated UTF-8",
			input: "ФЫвАф",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isUniq(tt.input)
			if got != tt.want {
				t.Errorf("isUniq(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}

}