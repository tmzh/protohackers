package main

import "testing"

func Test_isPrime(t *testing.T) {
	tests := []struct {
		name string
		n    float64
		want bool
	}{
		{"one", 1.0, false},
		{"two", 2.0, true},
		{"three", 3.0, true},
		{"four", 4.0, false},
		{"five", 5.0, true},
		{"seventeen", 17.0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.n); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
