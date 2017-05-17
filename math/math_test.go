package mathm

import (
	"testing"
)

var x int

func TestAbs(t *testing.T) {
	if value := Abs(1); value != 1 {
		t.Errorf("expected 1 got %v", value)
	}

	if value := Abs(-1); value != 1 {
		t.Errorf("expected 1 got %v", value)
	}
}

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = Abs(1 - 2*(i&1))
	}
}

func BenchmarkAbsPositive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = Abs(1)
	}
}

func BenchmarkAbsNegative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = Abs(-1)
	}
}

func TestGcd(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 1, 1},
		{1836311903, 1134903170, 1}, // two big Fibonacci's numbers
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Gcd(tt.a, tt.b); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLcm(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 1, 1},
		{1836311903, 1134903170, 1}, // two big Fibonacci's numbers
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Lcm(tt.a, tt.b); got != tt.want {
				t.Errorf("Lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}
