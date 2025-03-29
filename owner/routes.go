package owner

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Class represents a class that a studio owner can create
type Class struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Capacity  int       `json:"capacity"`
}

// for now maintaining as global variable but in real world it should be in db/per user level settings
var Classes []Class
var classIDCounter int

func RegisterRoutes(router *mux.Router) *mux.Router {
	ownerRouter := router.PathPrefix("/owner").Subrouter()
	ownerRouter.HandleFunc("/classes", CreateClassHandler).Methods("POST")
	return ownerRouter
}

// CreateClassHandler handles the creation of a new class
func CreateClassHandler(w http.ResponseWriter, r *http.Request) {

	// read the request body
	var newClass Class
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newClass)
	if err != nil {
		http.Error(w, "Invalid request body"+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if newClass.Name == "" || newClass.StartDate.IsZero() || newClass.EndDate.IsZero() || newClass.Capacity <= 0 {
		http.Error(w, "Missing or invalid class details", http.StatusBadRequest)
		return
	}

	// Check if the booking date is within the class date range
	if newClass.StartDate.After(newClass.EndDate) || newClass.EndDate.Before(newClass.StartDate) {
		http.Error(w, "Booking date is outside of the class range", http.StatusBadRequest)
		return
	}

	// Increment class ID for uniqueness
	classIDCounter++
	newClass.ID = classIDCounter

	// Save the class to memory
	Classes = append(Classes, newClass)

	// Respond with the created class
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newClass)
}
