package member

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ttanuja01/abc-glofox-api/owner"
)

// Booking represents a member's booking for a class
type Booking struct {
	MemberName string    `json:"name"`
	ClassDate  time.Time `json:"date"`
	ClassID    int       `json:"class_id"`
}

var bookings []Booking
var bookingIDCounter int

func RegisterRoutes(router *mux.Router) *mux.Router {
	memberRouter := router.PathPrefix("/member").Subrouter()
	memberRouter.HandleFunc("/bookings", BookClassHandler).Methods("POST")
	memberRouter.HandleFunc("/classes", GetAvailableClasses).Methods("GET")
	return memberRouter
}

// BookClassHandler handles the booking of a class
func BookClassHandler(w http.ResponseWriter, r *http.Request) {
	var newBooking Booking
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newBooking)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if newBooking.MemberName == "" || newBooking.ClassDate.IsZero() || newBooking.ClassID <= 0 {
		http.Error(w, "Missing or invalid booking details", http.StatusBadRequest)
		return
	}

	// Check if the class exists for the given class ID
	var foundClass *owner.Class
	for _, c := range owner.Classes {
		if c.ID == newBooking.ClassID {
			foundClass = &c
			break
		}
	}

	if foundClass == nil {
		http.Error(w, "Class not found", http.StatusNotFound)
		return
	}

	// Check if the booking date is within the class date range
	if newBooking.ClassDate.Before(foundClass.StartDate) || newBooking.ClassDate.After(foundClass.EndDate) {
		http.Error(w, "Booking date is outside of the class range", http.StatusBadRequest)
		return
	}

	// Increment booking ID for uniqueness
	bookingIDCounter++
	newBooking.ClassDate = newBooking.ClassDate

	// Save the booking
	bookings = append(bookings, newBooking)

	// Respond with the booking details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBooking)
}

// GetAvailableClasses returns the list of available classes
func GetAvailableClasses(w http.ResponseWriter, r *http.Request) {
	// Respond with the booking details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(owner.Classes)

}
