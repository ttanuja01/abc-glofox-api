package member

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/ttanuja01/abc-glofox-api/owner"
)

func init() {
	// Initialize some mock classes in the owner module (you can skip this part if you already have a mock in place)
	owner.Classes = []owner.Class{
		{
			ID:        1,
			Name:      "Yoga Class",
			StartDate: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 4, 1, 11, 0, 0, 0, time.UTC),
			Capacity:  20,
		},
		{
			ID:        2,
			Name:      "Pilates Class",
			StartDate: time.Date(2025, 4, 2, 10, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 4, 2, 11, 0, 0, 0, time.UTC),
			Capacity:  15,
		},
	}
}

func TestBookClass(t *testing.T) {
	// Test case: Valid booking
	t.Run("Valid Booking", func(t *testing.T) {
		// Setup request body
		bookingRequest := Booking{
			MemberName: "John Doe",
			ClassDate:  time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			ClassID:    1,
		}
		body, _ := json.Marshal(bookingRequest)

		req, err := http.NewRequest("POST", "/member/bookings", bytes.NewBuffer(body))
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
		var bookingResponse Booking
		if err := json.NewDecoder(rr.Body).Decode(&bookingResponse); err != nil {
			t.Fatal(err)
		}

		if bookingResponse.MemberName != "John Doe" {
			t.Errorf("Expected member name 'John Doe', got %s", bookingResponse.MemberName)
		}
	})

	// Test case: Invalid booking (missing required fields)
	t.Run("Invalid Booking - Missing Fields", func(t *testing.T) {
		// Invalid booking (missing class ID)
		bookingRequest := Booking{
			MemberName: "John Doe",
			ClassDate:  time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
		}
		body, _ := json.Marshal(bookingRequest)

		req, err := http.NewRequest("POST", "/member/bookings", bytes.NewBuffer(body))
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
	})

	// Test case: Invalid booking (class not found)
	t.Run("Invalid Booking - Class Not Found", func(t *testing.T) {
		// Invalid class ID
		bookingRequest := Booking{
			MemberName: "John Doe",
			ClassDate:  time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			ClassID:    999, // Non-existent class ID
		}
		body, _ := json.Marshal(bookingRequest)

		req, err := http.NewRequest("POST", "/member/bookings", bytes.NewBuffer(body))
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
		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("Expected status code %d, got %d", http.StatusNotFound, status)
		}
	})

	// Test case: Booking date outside the class range
	t.Run("Invalid Booking - Date Outside Class Range", func(t *testing.T) {
		// Booking outside of class range (class ends on April 1st, trying to book for April 2nd)
		bookingRequest := Booking{
			MemberName: "John Doe",
			ClassDate:  time.Date(2025, 4, 2, 10, 0, 0, 0, time.UTC),
			ClassID:    1,
		}
		body, _ := json.Marshal(bookingRequest)

		req, err := http.NewRequest("POST", "/member/bookings", bytes.NewBuffer(body))
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
	})
}

func TestGetAvailableClasses(t *testing.T) {
	// Test case: Get all available classes
	t.Run("Get Available Classes", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/member/classes", nil)
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
		var classes []owner.Class
		if err := json.NewDecoder(rr.Body).Decode(&classes); err != nil {
			t.Fatal(err)
		}

		if len(classes) != len(owner.Classes) {
			t.Errorf("Expected %d classes, got %d", len(owner.Classes), len(classes))
		}
	})
}
