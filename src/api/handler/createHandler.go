package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
	"tubes_3_dnAAAAA/regex"
	"tubes_3_dnAAAAA/stringmatching"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Struct untuk menyimpan data penyakit yang dibuat
type PenyakitFields struct {
	Nama_penyakit string
	Sequence_dna  string
}

// Struct untuk menyimpan data result dari respons
// yang akan diberikan ketika melakukan pembuatan data penyakit
type PenyakitResult struct {
	Message string
	Data    PenyakitFields
}

// Handler yang digunakan untuk membuat penyakit baru
func CreatePenyakit(response http.ResponseWriter, request *http.Request) {
	// Mengambil body, berupa data namaPenyakit dan sequenceDNA yang tersimpan dalam file,
	// dari request dengan ukuran maksimal 10 * 1024 * 1024 bytes
	request.ParseMultipartForm(10 * 1024 * 1024)
	file, _, err := request.FormFile("sequenceDNA")
	namaPenyakit := request.FormValue("namaPenyakit")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// Membaca sequence DNA dan mengecek apakah sequence DNA yang dibaca valid atau tidak
	isValid, dnaSequence := regex.ValidateDNASequence(file)

	// Jika tidak valid akan langsung mereturn respons berupa pesan error "Sequence DNA tidak valid"
	response.Header().Set("Content-Type", "application/json")

	if !isValid {
		response.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(response).Encode(PenyakitResult{
			Message: "Sequence DNA tidak valid",
		})
		return
	}

	// Jika valid maka data penyakit akan dibuat dan disimpan ke database
	const uri = "mongodb+srv://tubes-stima-dnAAAAA:tubes-stima-dnAAAAA@cluster0.lis3q.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	col := client.Database("dna").Collection("penyakit")
	fmt.Println("Successfully created collection. ", reflect.TypeOf(col))

	oneDoc := PenyakitFields{
		Nama_penyakit: namaPenyakit,
		Sequence_dna:  dnaSequence,
	}

	result, insertErr := col.InsertOne(context.TODO(), oneDoc)
	if insertErr != nil {
		fmt.Println(insertErr)
	} else {
		fmt.Println("Inserted a single document: ", reflect.TypeOf(result))
	}

	// Mengirimkan respons berupa pesan sukses dan data penyakit yang dibuat
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(PenyakitResult{
		Message: "Successfully created penyakit",
		Data:    oneDoc,
	})
}

// Struct untuk menyimpan data tes DNA yang dibuat
type TesDNAFields struct {
	Nama_pengguna        string
	Sequence_dna         string
	Prediksi_penyakit    string
	Tanggal              string
	Hasil_tes            bool
	Persentase_kemiripan float64
}

// Struct untuk menyimpan data result dari respons
// yang akan diberikan ketika melakukan pembuatan tes DNA
type TesDNAResult struct {
	Message            string
	Data               TesDNAFields
	StartingIndexFound int
	ComparisonCount    int
}

func CreateTesDNA(response http.ResponseWriter, request *http.Request) {
	// Mengambil body, berupa data namaPegguna, prediksiPenyakit, dan sequenceDNA
	// yang tersimpan dalam file, dari request dengan ukuran maksimal 10 * 1024 * 1024 bytes
	request.ParseMultipartForm(10 * 1024 * 1024)
	file, _, err := request.FormFile("sequenceDNA")
	namaPengguna := request.FormValue("namaPengguna")
	prediksiPenyakit := request.FormValue("prediksiPenyakit")
	metodeStringMatching := request.FormValue("metodeStringMatching")

	tanggal := time.Now().Format("02 January 2006")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// Membaca sequence DNA dan mengecek apakah sequence DNA yang dibaca valid atau tidak
	isValid, dnaSequencePasien := regex.ValidateDNASequence(file)

	// Jika tidak valid akan langsung mereturn respons berupa pesan error "Sequence DNA pasien tidak valid"
	response.Header().Set("Content-Type", "application/json")
	if !isValid {
		response.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(response).Encode(TesDNAResult{
			Message: "Sequence DNA pasien tidak valid",
		})
		return
	}

	// Jika valid maka data tes DNA akan dibuat dan disimpan ke database
	const uri = "mongodb+srv://tubes-stima-dnAAAAA:tubes-stima-dnAAAAA@cluster0.lis3q.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	tesDNACollections := client.Database("dna").Collection("tes_dna")
	penyakitCollections := client.Database("dna").Collection("penyakit")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Mengambil data penyakit yang sesuai dengan prediksi penyakit
	var penyakitBson bson.M
	if err := penyakitCollections.FindOne(ctx, bson.M{"nama_penyakit": prediksiPenyakit}).Decode(&penyakitBson); err != nil {
		// Apabila tidak terdapat data penyakit maka akan langsung
		// mereturn respons berupa pesan error "Data penyakit tidak ditemukan"
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(TesDNAResult{
			Message: "Data penyakit tidak ditemukan",
		})
		return
	}

	// Sequence DNA penyakit akan diambil dari data penyakit yang ditemukan
	dnaSequencePenyakit := penyakitBson["sequence_dna"].(string)

	var isMatch bool
	var startingIndexFound int
	var comparisonCount int
	var matchPercentage float64

	// Memeriksa apakah sequence DNA pasien sama dengan sequence DNA penyakit

	// Menggunakan metode KMP atau BM
	if metodeStringMatching == "KMP" {
		isMatch, startingIndexFound, comparisonCount = stringmatching.StringMatching(dnaSequencePenyakit, dnaSequencePasien)
		matchPercentage = 1
	} else if metodeStringMatching == "BM" {
		isMatch, startingIndexFound, comparisonCount = stringmatching.BmStringMatching(dnaSequencePenyakit, dnaSequencePasien)
		matchPercentage = 1
	}

	// Jika sequence DNA pasien dan penyakit tidak sama
	// maka akan ditentukan persentase kemiripannya
	if !isMatch {
		matchPercentage = stringmatching.SequenceSimilarity(dnaSequencePenyakit, dnaSequencePasien)

		if matchPercentage > 0.8 {
			isMatch = true
		}
	}

	fmt.Println("Successfully created collection. ", reflect.TypeOf(tesDNACollections))
	oneDoc := TesDNAFields{
		Nama_pengguna:        namaPengguna,
		Sequence_dna:         dnaSequencePasien,
		Prediksi_penyakit:    prediksiPenyakit,
		Tanggal:              tanggal,
		Hasil_tes:            isMatch,
		Persentase_kemiripan: matchPercentage,
	}

	result, insertErr := tesDNACollections.InsertOne(context.TODO(), oneDoc)
	if insertErr != nil {
		fmt.Println(insertErr)
	} else {
		fmt.Println("Inserted a single document: ", reflect.TypeOf(result))
	}

	// Mengirimkan respons berupa pesan sukses dan data penyakit yang dibuat
	response.WriteHeader(http.StatusCreated)

	json.NewEncoder(response).Encode(TesDNAResult{
		Message:            "Successfully created tes DNA",
		Data:               oneDoc,
		StartingIndexFound: startingIndexFound,
		ComparisonCount:    comparisonCount,
	})
}
