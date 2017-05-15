package mathm

import "testing"

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
