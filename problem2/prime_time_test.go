package main

import "testing"

func Test_isPrime(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"one", 1, false},
		{"two", 2, true},
		{"three", 3, true},
		{"four", 4, false},
		{"five", 5, true},
		{"seventeen", 17, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.n); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
