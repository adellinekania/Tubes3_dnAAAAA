package regex

import (
	"strconv"
	"strings"
)

func ReadQuery(query string) (bool, string, string, string) {
	arrQuery := strings.Split(query, "-")

	_, isNotDate := strconv.Atoi(arrQuery[0])

	date := ""
	penyakit := ""

	// Kalo tanggal (angka depannya)
	if isNotDate == nil {
		if len(arrQuery) < 3 {
			return false, "Format tanggal salah", date, penyakit
		}
		if _, isNotMonth := strconv.Atoi(arrQuery[1]); isNotMonth == nil {
			return false, "Format bulan salah, tidak boleh angka", date, penyakit
		}
		if !(arrQuery[1] == "January" || arrQuery[1] == "February" || arrQuery[1] == "March" || arrQuery[1] == "April" || arrQuery[1] == "May" || arrQuery[1] == "June" || arrQuery[1] == "July" || arrQuery[1] == "August" || arrQuery[1] == "September" || arrQuery[1] == "October" || arrQuery[1] == "November" || arrQuery[1] == "December") {
			return false, "Format bulan salah, bulan tidak valid", date, penyakit
		}
		if _, isNotYear := strconv.Atoi(arrQuery[2]); isNotYear != nil {
			return false, "Format tahun salah, harus angka", date, penyakit
		}

		date += arrQuery[0] + " " + arrQuery[1] + " " + arrQuery[2]

		if len(arrQuery) > 3 {
			for i := 3; i < len(arrQuery)-1; i++ {
				penyakit += arrQuery[i] + " "
			}
			penyakit += arrQuery[len(arrQuery)-1]
		}

		return true, "Data Benar", date, penyakit
	} else {
		for i := 0; i < len(arrQuery)-1; i++ {
			penyakit += arrQuery[i] + " "
		}
		penyakit += arrQuery[len(arrQuery)-1]
		return true, "Query Benar", date, penyakit
	}
}
