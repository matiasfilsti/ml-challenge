package main

import (
	"fmt"
	"log"
	"ml-challenge/modules/helper"
	"ml-challenge/modules/metrics"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const apiUrl string = "https://api.mercadolibre.com"

type Response struct {
	Message string `json:"message"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func categories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	urlResponse := helper.UrlGet(fmt.Sprintf("%v/categories/%v", apiUrl, id))
	metrics.RecordMetrics()
	fmt.Println(id)
	fmt.Fprint(w, string(urlResponse))
}

func items(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	urlResponse := helper.UrlGet(fmt.Sprintf("%v/items/%v", apiUrl, id))
	metrics.RecordMetrics()
	fmt.Println(id)
	fmt.Fprint(w, string(urlResponse))
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/categories/{id}", categories)
	r.HandleFunc("/items/{id}", items)
	//http.Handle("/metrics", promhttp.Handler())
	r.Path("/metrics").Handler(promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {

	handleRequests()
}
