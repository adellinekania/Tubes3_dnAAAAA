package stringmatching

import (
	"fmt"
)

func BorderFunction(inputPattern string) []int {
	inputPatternLength := int(len(inputPattern))
	b := make([]int, inputPatternLength)
	// Asumsi panjang pattern selalu lebih dari 2
	b[0] = 0
	b[1] = 0
	for i := 2; i < inputPatternLength; i++ {
		pattern := string(inputPattern[0:i])
		// log.Println("PATTERN", pattern)
		maxLength := 0
		patternLength := len(pattern)
		for j := 0; j < patternLength-1; j++ {
			prefix := string(pattern[0:(j + 1)])
			sufix := string(pattern[(patternLength - j - 1):patternLength])
			if prefix == sufix && j+1 > maxLength {
				maxLength = j + 1
			}
		}
		b[i] = maxLength
	}
	return b
}

func StringMatching(inputPattern string, inputText string) (bool, int, int) {
	b := BorderFunction(inputPattern)
	lenInputPattern := len(inputPattern)
	lenInputText := len(inputText)
	i := 0
	j := 0
	isMatch := false
	count := 1
	startingIndex := 0
	jumlahPerbandingan := 0
	for !isMatch && i < lenInputText-lenInputPattern {
		bValue := b[j]
		if count > 1 && j == 0 {
			i++
		}
		startingIndex = i - bValue
		j = bValue

		// fmt.Println("Perhitungan ke:", count)
		// fmt.Println("startingIndex:", startingIndex)
		// fmt.Println("i:", i)
		// fmt.Println("j:", j)

		isMatch = true
		for isMatch && j < lenInputPattern {
			jumlahPerbandingan++
			if inputPattern[j] != inputText[i] {
				isMatch = false
			} else {
				i++
				j++
			}
		}
		count++
	}

	fmt.Println(isMatch, startingIndex, "Jumlah perbandungan", jumlahPerbandingan)
	return isMatch, startingIndex, jumlahPerbandingan
}
