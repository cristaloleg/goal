package string

// EditDist computes distance between strings
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
