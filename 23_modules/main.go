package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello modules..")
	//? go to the go path then to \pkg\mod there you will get cache where all the downloads are getting stored
	//! gorilla mux path - cache\download\github.com\gorilla
	greet()
	// r := mux.NewRouter()
	// r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4500", r))
}

func greet() {
	fmt.Println("Hello this is user routes!!!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to akhil's learning of golang</h1>"))
}
