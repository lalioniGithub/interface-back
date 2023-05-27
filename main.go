package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Solicitor struct {
	ID       int    `json:"ID"`
	Dni      string `json:"dni"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Role     string `json:"role"`
}

type AllSolicitors []Solicitor

var solicitors = AllSolicitors{
	{
		ID:       1,
		Dni:      "39419524",
		Name:     "Leonardo",
		LastName: "Alioni",
		Role:     "Tester",
	},
}

func getSolicitors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solicitors)
}

func createSolicitor(w http.ResponseWriter, r *http.Request) {
	var newSolicitor Solicitor

	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Invalid request.\n%s", err)
	}

	json.Unmarshal(reqBody, &newSolicitor)

	newSolicitor.ID = len(solicitors) + 1

	solicitors = append(solicitors, newSolicitor)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newSolicitor)
	fmt.Fprintf(w, "(%s %s - ID: %d) has been successfully added from solicitors.", newSolicitor.Name, newSolicitor.LastName, newSolicitor.ID)
}

func getSolicitor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	solicitorId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid request. ID must be of type int64. \n%s", err)
		return
	}

	for _, solicitor := range solicitors {
		if solicitor.ID == solicitorId {
			json.NewEncoder(w).Encode(solicitor)
		}
	}
}

func deleteSolicitor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	solicitorId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid request. ID must be of type int64. \n%s", err)
		return
	}

	for i, solicitor := range solicitors {
		if solicitor.ID == solicitorId {
			solicitors = append(solicitors[:i], solicitors[i + 1:]...)
			fmt.Fprintf(w, "(%s %s - ID: %d) has been successfully removed from solicitors.", solicitor.Name, solicitor.LastName, solicitor.ID)
		}
	}
}

func updateSolicitor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	solicitorId, err := strconv.Atoi(vars["id"])

	var updatedSolicitor Solicitor

	if err != nil {
		fmt.Fprintf(w, "Invalid request. ID must be of type int64. \n%s", err)
		return
	}

	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Request data invalid. \n%s", err)
	}

	json.Unmarshal(reqBody, &updatedSolicitor)

	for i, s := range solicitors {
		if s.ID == solicitorId {
			w.Header().Set("Content-Type", "application/json")
			solicitors = append(solicitors[:i], solicitors[i+1:]...)
			updatedSolicitor.ID = solicitorId
			solicitors = append(solicitors, updatedSolicitor)
			fmt.Fprintf(w, "(%s %s - ID: %d) has been successfully updated.", s.Name, s.LastName, s.ID)
		}
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my REST API")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/getsolicitors", getSolicitors).Methods("GET")
	router.HandleFunc("/getsolicitor/{id}", getSolicitor).Methods("GET")
	router.HandleFunc("/createsolicitor", createSolicitor).Methods("POST")
	router.HandleFunc("/removesolicitor/{id}", deleteSolicitor).Methods("DELETE")
	router.HandleFunc("/updatesolicitor/{id}", updateSolicitor).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}
