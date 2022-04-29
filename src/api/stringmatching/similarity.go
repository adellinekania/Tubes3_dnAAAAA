package stringmatching

import "math"

func LongestCommonSubsequence(inputPattern string, inputText string) int {
	m := len(inputPattern)
	n := len(inputText)
	Table := make([][]int, m+1)
	for i := range Table {
		Table[i] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 || j == 0 {
				Table[i][j] = 0
			} else if inputPattern[i-1] == inputText[j-1] {
				Table[i][j] = Table[i-1][j-1] + 1
			} else {
				Table[i][j] = int(math.Max(float64(Table[i-1][j]), float64(Table[i][j-1])))
			}
		}
	}

	return Table[m][n]
}

func SequenceSimilarity(inputPattern string, inputText string) float64 {
	lcs := LongestCommonSubsequence(inputPattern, inputText)
	return float64(lcs) / math.Max(float64(len(inputPattern)), float64(len(inputText)))
}
