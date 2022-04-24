package regex

import (
	"bufio"
	"mime/multipart"
	"regexp"
)

func ValidateDNASequence(f multipart.File) (bool, string) {
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	var dnaSequence string
	for scanner.Scan() {
		// do something with a line
		dnaSequence += scanner.Text()
	}
	match, _ := regexp.MatchString("^[CAGT]+$", dnaSequence)
	if !match {
		return false, ""
	}

	return true, dnaSequence
}
