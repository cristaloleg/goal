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

func TestEditDistEx(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{"dog", "bug", 2},
		{"kitten", "sitting", 3},
		{"saturday", "sunday", 3},

		{"xca", "xabc", 2},
		{"cax", "abcx", 2},
		{"ca", "bc", 2},
		{"hello", "hello", 0},
		{"hello", "helo", 1},
		{"hello", "jello", 1},
		{"hello", "helol", 1},
		{"hello", "hellol", 1},
		{"hello", "heloll", 2},
		{"hello", "cheese", 4},
		{"hello", "saint", 5},
		{"hallo", "saint", 4},
		{"hello", "", 5},
		{"smtih", "smith", 1},
		{"snapple", "apple", 2},
		{"testing", "testtn", 2},
		{"saturday", "sunday", 3},
		{"Saturday", "saturday", 1},
		{"orange", "pumpkin", 7},
		{"gifts", "profit", 5},
		// {"ca", "ac", 1},
		// {"ca", "abc", 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := EditDistEx(tt.a, tt.b); got != tt.want {
				t.Errorf("got %v want %v :: `%v` & `%v`", got, tt.want, tt.a, tt.b)
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
