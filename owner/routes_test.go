package owner

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	// Initialize any global variables if needed (e.g., resetting the Classes and classIDCounter)
	Classes = []Class{}
	classIDCounter = 0
}

func TestCreateClassHandler(t *testing.T) {
	// Test case: Valid class creation
	t.Run("Valid Class Creation", func(t *testing.T) {
		// Prepare request body
		classRequest := Class{
			Name:      "Yoga Class",
			StartDate: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 5, 1, 11, 0, 0, 0, time.UTC),
			Capacity:  20,
		}
		body, _ := json.Marshal(classRequest)

		req, err := http.NewRequest("POST", "/owner/classes", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		// Mock the router
		router := mux.NewRouter()
		RegisterRoutes(router)

		// Create a recorder to capture the response
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, status)
		}

		// Check the response body
		var createdClass Class
		if err := json.NewDecoder(rr.Body).Decode(&createdClass); err != nil {
			t.Fatal(err)
		}

		// Validate the response body matches the class request
		if createdClass.Name != classRequest.Name {
			t.Errorf("Expected class name %s, got %s", classRequest.Name, createdClass.Name)
		}
		if !createdClass.StartDate.Equal(classRequest.StartDate) {
			t.Errorf("Expected start date %v, got %v", classRequest.StartDate, createdClass.StartDate)
		}
		if !createdClass.EndDate.Equal(classRequest.EndDate) {
			t.Errorf("Expected end date %v, got %v", classRequest.EndDate, createdClass.EndDate)
		}
		if createdClass.Capacity != classRequest.Capacity {
			t.Errorf("Expected capacity %d, got %d", classRequest.Capacity, createdClass.Capacity)
		}
	})

	// Test case: Invalid class creation (missing required fields)
	t.Run("Invalid Class Creation - Missing Fields", func(t *testing.T) {
		// Invalid request (missing the Name field)
		classRequest := Class{
			StartDate: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 4, 1, 11, 0, 0, 0, time.UTC),
			Capacity:  20,
		}
		body, _ := json.Marshal(classRequest)

		req, err := http.NewRequest("POST", "/owner/classes", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		// Mock the router
		router := mux.NewRouter()
		RegisterRoutes(router)

		// Create a recorder to capture the response
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, status)
		}

		// Check the response body
		if rr.Body.String() != "Missing or invalid class details\n" {
			t.Errorf("Expected error message 'Missing or invalid class details', got %s", rr.Body.String())
		}
	})

	// Test case: Invalid class creation (invalid date range)
	t.Run("Invalid Class Creation - Invalid Date Range", func(t *testing.T) {
		// Invalid request (end date is before start date)
		classRequest := Class{
			Name:      "Yoga Class",
			StartDate: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 3, 31, 11, 0, 0, 0, time.UTC), // End date is before start date
			Capacity:  20,
		}
		body, _ := json.Marshal(classRequest)

		req, err := http.NewRequest("POST", "/owner/classes", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		// Mock the router
		router := mux.NewRouter()
		RegisterRoutes(router)

		// Create a recorder to capture the response
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, status)
		}

		// Check the response body
		if rr.Body.String() != "Booking date is outside of the class range\n" {
			t.Errorf("Expected error message 'Booking date is outside of the class range', got %s", rr.Body.String())
		}
	})
}
