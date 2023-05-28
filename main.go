package main

import (
	"log"
	"net/http"


	"github.com/gorilla/mux"
	"github.com/lalioniGithub/interface-back/routes"
)





func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", routes.IndexRoute)
	router.HandleFunc("/getsolicitors", routes.GetSolicitors).Methods("GET")
	router.HandleFunc("/getsolicitor/{id}", routes.GetSolicitor).Methods("GET")
	router.HandleFunc("/createsolicitor", routes.CreateSolicitor).Methods("POST")
	router.HandleFunc("/removesolicitor/{id}", routes.DeleteSolicitor).Methods("DELETE")
	router.HandleFunc("/updatesolicitor/{id}", routes.UpdateSolicitor).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}
