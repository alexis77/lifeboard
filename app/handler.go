package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/articles", getArticles).Methods("GET")
	router.HandleFunc("/articles", createArticle).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index endpoint")
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	//fmt.Fprintf(article)
	list = append(list, article)

	json.NewEncoder(w).Encode(article)
}
