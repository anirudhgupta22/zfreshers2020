package main

import (
	"encoding/json"
	"github.com/spf13/cast"
	"log"
	"net/http"
)

type Id int32

type Zoman struct {
	Id      Id
	Name    string
	Age     int32
	College string
	Team    string
}

type ZomanDetailsResponse struct {
	Name    string `json:"name"`
	Team    string `json"team"`
	College string `json:"college"`
}

//in memory database
var zomans map[Id]Zoman

func init() {
	zomans = map[Id]Zoman{
		1: {
			Id:      1,
			Name:    "Tom",
			Age:     23,
			College: "IIT",
			Team:    "Reviews",
		},
		2: {
			Id:      1,
			Name:    "John",
			Age:     21,
			College: "NIT",
			Team:    "Merchant",
		},
		3: {
			Id:      3,
			Name:    "Sara",
			Age:     20,
			College: "IIIT",
			Team:    "Search",
		},
		4: {
			Id:      4,
			Name:    "Harry",
			Age:     23,
			College: "HBTI",
			Team:    "SRE",
		},
	}
}

func getZomanDetails(id Id) (Zoman, bool) {
	z, ok := zomans[id]
	return z, ok
}

func handleZomansApi(w http.ResponseWriter, r *http.Request) {
	log.Printf("New request on %s", r.URL)
	id := r.FormValue("id")
	details, found := getZomanDetails(Id(cast.ToInt64(id)))
	if !found {
		http.Error(w, "Invalid id", http.StatusInternalServerError)
	} else {
		response := ZomanDetailsResponse{details.Name, details.Team, details.College}
		json.NewEncoder(w).Encode(response)
	}

}
