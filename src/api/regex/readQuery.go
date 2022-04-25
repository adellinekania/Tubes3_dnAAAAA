package regex

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func splitWord(word string) []string {
	array := regexp.MustCompile("[\\-\\/\\.\\s]+").Split(word, -1)
	return array
}

func ReadQuery(query string) (bool, string, string, string) {
	months := [...]string{
		"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December",
	}
	valid := false
	validQueryRegex := [...]string{
		"\\d{2}[/.-]\\d{2}[/.-]\\d{4}.*",
		"\\d{2} \\d{2} \\d{4}.*",
		"\\d{2} (January|February|March|April|May|June|July|August|September|October|November|December) \\d{4}.*",
	}
	i := 0

	_, isNotDate := strconv.Atoi(string(query[0]))

	if isNotDate != nil {
		return true, "Query Benar", "", query
	}

	for i < len(validQueryRegex) && !valid {
		match, _ := regexp.MatchString(validQueryRegex[i], query)
		if match {
			valid = true
		} else {
			i++
		}
	}

	arrQuery := splitWord(query)
	if valid {
		m, _ := strconv.Atoi(arrQuery[1])
		if i != 2 {
			arrQuery[1] = months[m-1]
		}
		newDate := arrQuery[1] + " " + arrQuery[0] + ", " + arrQuery[2]
		_, err := time.Parse("January 02, 2006", newDate)
		if err != nil {
			return false, "Tanggal salah!", "", ""
		}
		return true, "Query Benar", newDate, strings.Join(arrQuery[3:], " ")
	} else {
		return false, "Format query salah!", "", ""
	}
}
