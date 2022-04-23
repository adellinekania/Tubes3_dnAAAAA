package stringmatching

import (
	"fmt"
	"math"
)

func BuildLast(inputPattern string) []int {
	last := make([]int, 128)

	for i := 0; i < 128; i++ {
		last[i] = -1
	}

	for i := 0; i < len(inputPattern); i++ {
		last[inputPattern[i]] = i
	}

	return last
}

func BmStringMatching(inputPattern string, inputText string) (bool, int, int) {
	fmt.Println("BM")
	last := BuildLast(inputPattern)
	n := len(inputText)
	m := len(inputPattern)
	i := m - 1
	j := m - 1
	jumlahPerbandingan := 0

	if i > n-1 {
		return false, 0, 0
	}

	for next := true; next; next = i <= n-1 {
		jumlahPerbandingan++
		if inputPattern[j] == inputText[i] {
			if j == 0 {
				return true, i, jumlahPerbandingan
			} else {
				i--
				j--
			}
		} else {
			lo := last[inputText[i]]
			i = i + m - int(math.Min(float64(j), float64(lo+1)))
			j = m - 1
		}
	}

	return false, 0, jumlahPerbandingan
}
