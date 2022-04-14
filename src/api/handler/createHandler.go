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

type PenyakitFields struct {
	Nama_penyakit string
	Sequence_dna  string
}

type PenyakitResult struct {
	Message string
	Data    PenyakitFields
}

func CreatePenyakit(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 * 1024 * 1024)
	file, _, err := request.FormFile("sequenceDNA")
	namaPenyakit := request.FormValue("namaPenyakit")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	isValid, dnaSequence := regex.ValidateDNASequence(file)

	response.Header().Set("Content-Type", "application/json")

	if !isValid {
		response.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(response).Encode(PenyakitResult{
			Message: "Penyakit ga valid",
		})
		return
	}

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

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(PenyakitResult{
		Message: "Successfully created penyakit",
		Data:    oneDoc,
	})
}

type TesDNAFields struct {
	Nama_pengguna     string
	Sequence_dna      string
	Prediksi_penyakit string
	Tanggal           string
	Hasil_tes         bool
}

type TesDNAResult struct {
	Message            string
	Data               TesDNAFields
	StartingIndexFound int
	ComparisonCount    int
}

func CreateTesDNA(response http.ResponseWriter, request *http.Request) {
	tanggal := time.Now().Format("02 January 2006")

	request.ParseMultipartForm(10 * 1024 * 1024)
	file, _, err := request.FormFile("sequenceDNA")
	namaPengguna := request.FormValue("namaPengguna")
	prediksiPenyakit := request.FormValue("prediksiPenyakit")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	isValid, dnaSequencePasien := regex.ValidateDNASequence(file)

	response.Header().Set("Content-Type", "application/json")
	if !isValid {
		response.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(response).Encode(TesDNAResult{
			Message: "DNA PAsien ga valid",
		})
		return
	}

	const uri = "mongodb+srv://tubes-stima-dnAAAAA:tubes-stima-dnAAAAA@cluster0.lis3q.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	tesDNACollections := client.Database("dna").Collection("tes_dna")
	penyakitCollections := client.Database("dna").Collection("penyakit")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var penyakitBson bson.M
	if err := penyakitCollections.FindOne(ctx, bson.M{"nama_penyakit": prediksiPenyakit}).Decode(&penyakitBson); err != nil {
		// File ga ketemu
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(TesDNAResult{
			Message: "Penyakit tidak ditemukan",
		})
		return
	}

	dnaSequencePenyakit := penyakitBson["sequence_dna"].(string)

	isMatch, startingIndexFound, comparisonCount := stringmatching.StringMatching(dnaSequencePenyakit, dnaSequencePasien)

	fmt.Println("Successfully created collection. ", reflect.TypeOf(tesDNACollections))
	oneDoc := TesDNAFields{
		Nama_pengguna:     namaPengguna,
		Sequence_dna:      dnaSequencePasien,
		Prediksi_penyakit: prediksiPenyakit,
		Tanggal:           tanggal,
		Hasil_tes:         isMatch,
	}

	result, insertErr := tesDNACollections.InsertOne(context.TODO(), oneDoc)
	if insertErr != nil {
		fmt.Println(insertErr)
	} else {
		fmt.Println("Inserted a single document: ", reflect.TypeOf(result))
	}

	response.WriteHeader(http.StatusCreated)

	json.NewEncoder(response).Encode(TesDNAResult{
		Message:            "Successfully created tes DNA",
		Data:               oneDoc,
		StartingIndexFound: startingIndexFound,
		ComparisonCount:    comparisonCount,
	})
}
