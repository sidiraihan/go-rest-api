package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//company struct
type Company struct {
	ID       string
	Name     string
	Email    string
	Ballance float64
}

var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test sslmode=disable password=")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	router.HandleFunc("/company", GetCompanies).Methods("GET")
	router.HandleFunc("/company/{id}", GetCompany).Methods("GET")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))
}

//get all companies
func GetCompanies(w http.ResponseWriter, r *http.Request) {
	var companies []Company
	db.Find(&companies)
	json.NewEncoder(w).Encode(&companies)
}

//get single companies
func GetCompany(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var companies Company
	db.First(&companies, params["id"])
	json.NewEncoder(w).Encode(&companies)
}
