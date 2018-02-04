package string

import (
	"bytes"
	"fmt"
)

// EditDist computes distance between strings
// Levenshtein algorithm
func EditDist(a, b string) int {
	len1, len2 := len(a), len(b)
	if len1 < len2 {
		return EditDist(b, a)
	}
	row1, row2 := make([]int, len2+1), make([]int, len2+1)

	for i := 0; i < len2+1; i++ {
		row2[i] = i
	}

	for i := 0; i < len1; i++ {
		row1[0] = i + 1

		for j := 0; j < len2; j++ {
			x := min(row2[j+1]+1, row1[j]+1)
			y := row2[j] + invBool2int(a[i] == b[j])
			row1[j+1] = min(x, y)
		}

		row1, row2 = row2, row1
	}
	return row2[len2]
}

// EditDistEx computes distance between strings
// Damerau-Levenshtein algorithm
func EditDistEx(a, b string) int {
	len1, len2 := len(a), len(b)
	if len1 == 0 {
		return len2
	}
	if len2 == 0 {
		return len1
	}
	if len1 < len2 {
		return EditDistEx(b, a)
	}
	curr, next := 0, 0
	row := make([]int, len2+1)

	for i := 0; i < len2+1; i++ {
		row[i] = i
	}

	for i := 0; i < len1; i++ {
		curr = i + 1

		for j := 0; j < len2; j++ {
			cost := invBool2int(a[i] == b[j] || (i > 0 && j > 0 && a[i-1] == b[j] && a[i] == b[j-1]))
			fmt.Printf("%v %v == %v\n", a[i], b[j], cost)

			next = min(min(
				row[j+1]+1,
				row[j]+cost),
				curr+1)

			row[j], curr = curr, next
		}
		row[len2] = next
		fmt.Printf("%v\n", row)
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
	return next
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func invBool2int(b bool) int {
	if !b {
		return 1
	}
	return 0
}

// primeRK is the prime base used in Rabin-Karp algorithm.
const primeRK = 16777619

// SearchRabinKarp search `sep` in `s` using Rabin-Karp algorithm
// based on src/bytes/bytes.go
func SearchRabinKarp(s, sep []byte) int {
	hashsep, pow := hashStr(sep)
	n := len(sep)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	if h == hashsep && bytes.Equal(s[:n], sep) {
		return 0
	}
	for i := n; i < len(s); {
		h *= primeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashsep && bytes.Equal(s[i-n:i], sep) {
			return i - n
		}
	}
	return -1
}

// hashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func hashStr(sep []byte) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}
