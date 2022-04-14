package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"tubes_3_dnAAAAA/regex"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://tubes-stima-dnAAAAA:tubes-stima-dnAAAAA@cluster0.lis3q.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

type TesDNAResults struct {
	Message string
	Data    []bson.M
}

type Tes struct {
    Message string
}

func ApiTest(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusCreated)

	json.NewEncoder(response).Encode(Tes{
		Message: "Successfully get tes DNA",
	})

}

func GetAllTesDNA(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	query := vars["query"]
	fmt.Println("INI WUERY", query == "")

	isQueryValid, message, date, penyakit := true, "", "", ""

	if query != "" {
		isQueryValid, message, date, penyakit = regex.ReadQuery(query)
	}

	response.Header().Set("Content-Type", "application/json")
	if !isQueryValid {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(TesDNAResults{
			Message: message,
			Data:    nil,
		})
		return
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}

	defer client.Disconnect(ctx)

	tesDNACollection := client.Database("dna").Collection("tes_dna")

	cursor, err := tesDNACollection.Find(ctx, bson.M{})
	if date != "" && penyakit == "" {
		cursor, err = tesDNACollection.Find(ctx, bson.M{"tanggal": date})
	} else if penyakit != "" && date == "" {
		cursor, err = tesDNACollection.Find(ctx, bson.M{"prediksi_penyakit": penyakit})
	} else if penyakit != "" && date != "" {
		cursor, err = tesDNACollection.Find(ctx, bson.M{"tanggal": date, "prediksi_penyakit": penyakit})
	}

	if err != nil {
		fmt.Println(err)
	}

	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		fmt.Println(err)
	}

	response.WriteHeader(http.StatusCreated)

	json.NewEncoder(response).Encode(TesDNAResults{
		Message: "Successfully get tes DNA",
		Data:    result,
	})

}
