package main

import (
	"context"
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"tubes_3_dnAAAAA/handler"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb+srv://tubes-stima-dnAAAAA:tubes-stima-dnAAAAA@cluster0.lis3q.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

type spaHandler struct {
	staticPath string
	indexPath  string
}

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
	err = mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		return
	}
	router.HandleFunc("/api/upload", handler.CreatePenyakit).Methods("POST")
	router.HandleFunc("/api/tesDNA", handler.CreateTesDNA).Methods("POST")
	router.HandleFunc("/api/penyakit", handler.GetPenyakit).Methods("GET")
	router.HandleFunc("/api/tesDNA", handler.GetAllTesDNA).Methods("GET")
	router.HandleFunc("/api/tesDNA/{query}", handler.GetAllTesDNA).Methods("GET")

	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	distDir := filepath.Join(currentDir, "../client/dist")
	spa := spaHandler{staticPath: distDir, indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	router.HandleFunc("/api/test", handler.ApiTest).Methods("GET")
	err = http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)
	indexFile := filepath.Join(h.staticPath, h.indexPath)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, indexFile)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
