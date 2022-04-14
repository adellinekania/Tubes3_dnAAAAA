package regex

import (
	"bufio"
	"mime/multipart"
)

func ValidateDNASequence(f multipart.File) (bool, string) {
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	var dnaSequence string
	for scanner.Scan() {
		// do something with a line
		dnaSequence += scanner.Text()
	}

	for i := 0; i < len(dnaSequence); i++ {
		dna := string(dnaSequence[i])
		// fmt.Println(dna)
		if !(dna == "C" || dna == "G" || dna == "A" || dna == "T") {
			return false, ""
		}
	}

	return true, dnaSequence
}
