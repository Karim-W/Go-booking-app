package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"bookingapp.com/Controllers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/reserve", getReserves).Methods("GET")
	log.Fatal(http.ListenAndServe(":8022", r))
}

type shootme struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func getReserves(res http.ResponseWriter, req *http.Request) {
	a := shootme{FirstName: "ff", LastName: "hh"}
	fmt.Println(a.FirstName)
	Controllers.HandleGetReserves()
	res.Header().Set("content-type", "application/json")
	json.NewEncoder(res).Encode(a)
}
