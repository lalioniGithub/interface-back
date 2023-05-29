package routes

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/lalioniGithub/interface-back/db"
	"github.com/lalioniGithub/interface-back/models"
)

func GetSolicitors(w http.ResponseWriter, r *http.Request) {
	var solicitors []models.Solicitor
	db.DB.Find(&solicitors)

	json.NewEncoder(w).Encode(&solicitors)
}

func CreateSolicitor(w http.ResponseWriter, r *http.Request) {
	var newSolicitor models.Solicitor
	json.NewDecoder(r.Body).Decode(&newSolicitor)

	createdSolicitor := db.DB.Create(&newSolicitor)
	err := createdSolicitor.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&newSolicitor)

	fmt.Println("CORRECTO")
	w.WriteHeader(http.StatusOK)
}

func GetSolicitor(w http.ResponseWriter, r *http.Request) {
	var solicitor models.Solicitor
	params := mux.Vars(r)

	db.DB.First(&solicitor, params["id"])

	if solicitor.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Solicitor not found"))
		return
	}

	json.NewEncoder(w).Encode(&solicitor)
	w.WriteHeader(http.StatusOK)
}

func DeleteSolicitor(w http.ResponseWriter, r *http.Request) {
	var solicitor models.Solicitor
	params := mux.Vars(r)

	db.DB.First(&solicitor, params["id"])

	if solicitor.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Solicitor not found"))
		return
	}

	db.DB.Unscoped().Delete(&solicitor) //Con el "Unscoped()" lo borra físicamente
	w.WriteHeader(http.StatusOK)
}

func UpdateSolicitor(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// solicitorId, err := strconv.Atoi(vars["id"])

	// var updatedSolicitor models.Solicitor

	// if err != nil {
	// 	fmt.Fprintf(w, "Invalid request. ID must be of type int64. \n%s", err)
	// 	return
	// }

	// reqBody, err := io.ReadAll(r.Body)

	// if err != nil {
	// 	fmt.Fprintf(w, "Request data invalid. \n%s", err)
	// }

	// json.Unmarshal(reqBody, &updatedSolicitor)

	// for i, s := range solicitors {
	// 	if s.ID == solicitorId {
	// 		w.Header().Set("Content-Type", "application/json")
	// 		solicitors = append(solicitors[:i], solicitors[i+1:]...)
	// 		updatedSolicitor.ID = solicitorId
	// 		solicitors = append(solicitors, updatedSolicitor)
	// 		fmt.Fprintf(w, "(%s %s - ID: %d) has been successfully updated.", s.Name, s.LastName, s.ID)
	// 		break
	// 	}
	// }
}