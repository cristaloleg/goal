package string

import (
	"testing"
)

var xStringTest int

func TestEditDist(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{"dog", "bug", 2},
		{"kitten", "sitting", 3},
		{"saturday", "sunday", 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := EditDist(tt.a, tt.b); got != tt.want {
				t.Errorf("EditDist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkEditDist(t *testing.B) {
	a, b := "1234567890123456789", "123456789987654321"
	t.ReportAllocs()
	t.StopTimer()
	for i := 0; i < 10000; i++ {
		xStringTest = EditDist(a, b)
	}
	t.StartTimer()
}
