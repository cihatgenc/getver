package main

import (
	"encoding/json"
	//	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Return all MSSQL versions
func getVersions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	myobject, err := Crawler()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	if err := json.NewEncoder(w).Encode(myobject); err != nil {
		panic(err)
	}
}

// Return all MSSQL versions by Major version
func getVersionsByMajor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	mversion := vars["majorversion"]

	myobject, err := CrawlerMajor(mversion)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	if err := json.NewEncoder(w).Encode(myobject); err != nil {
		panic(err)
	}
}

// Return latest MSSQL version by Major version
func getlatestVersionByMajor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	mversion := vars["majorversion"]

	myobject, err := CrawlerMajorLatest(mversion)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	if err := json.NewEncoder(w).Encode(myobject); err != nil {
		panic(err)
	}
}
