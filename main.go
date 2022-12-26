package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Object struct {
	ID         string `json:"id"`
	IconUrl    string `json:"icon_url"`
	Url        string `json:"url"`
	Value      string `json:"value"`
	PetitionID string `json:"petition_id"`
}

const APIChuck = "https://api.chucknorris.io/jokes/random"

func fetchObject() (*Object, error) {
	println("Start")

	var objectReturn *Object

	objectFetching, err := http.Get(APIChuck)

	if err != nil {
		println("Error al optain data")
		return &Object{}, err
	}
	data, err := io.ReadAll(objectFetching.Body)

	if err != nil {
		println("Error al optener data")
		return &Object{}, err
	}
	if err := json.Unmarshal(data, &objectReturn); err != nil {
		fmt.Print("Unmarshal error")
	}

	return objectReturn, nil
}

func getObjects(w http.ResponseWriter, req *http.Request) {

	var arrObjectResponse []*Object

	for len(arrObjectResponse) < 25 {
		resp, err := fetchObject()
		if err != nil{
			println("Error to fetch data")
		}
		resp.PetitionID = uuid.New().String()
		arrObjectResponse = append(arrObjectResponse, resp)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(arrObjectResponse)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getObjects).Methods("GET")

	log.Fatal(http.ListenAndServe(":3030",router))
}
