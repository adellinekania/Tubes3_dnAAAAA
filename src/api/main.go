package main

import (
	"context"
	"fmt"
	"net/http"
	"tubes_3_dnAAAAA/handler"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb+srv://tubes-stima-dnAAAAA:tubes-stima-dnAAAAA@cluster0.lis3q.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

func main() {

	// Melakukan koneksi ke mongoose database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	// Membuat router
	router := mux.NewRouter()
	router.HandleFunc("/api/upload", handler.CreatePenyakit).Methods("POST")
	router.HandleFunc("/api/tesDNA", handler.CreateTesDNA).Methods("POST")
	router.HandleFunc("/api/penyakit", handler.GetPenyakit).Methods("GET")
	router.HandleFunc("/api/tesDNA", handler.GetAllTesDNA).Methods("GET")
	router.HandleFunc("/api/tesDNA/{query}", handler.GetAllTesDNA).Methods("GET")
	router.HandleFunc("/api/test", handler.ApiTest).Methods("GET")
	err = http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
