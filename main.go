package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ttanuja01/abc-glofox-api/member"
	"github.com/ttanuja01/abc-glofox-api/owner"
)

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Register the routes for owner and member modules
	owner.RegisterRoutes(r)
	member.RegisterRoutes(r)

	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}


//This repo has two modules: owner and member. The owner module allows studio owners to create classes, while the member module allows members to book classes. 
// The code uses the Gorilla Mux router for handling HTTP requests and responses. 
// How to test the API:
//1.mke sure you have Go installed on your machine. You can download it from https://golang.org/dl/ and run go.mod tidy to install the dependencies.
//2. run the code using the command go run main.go. This will start a local server on port 8000.
//3.install thunder client or postman to test the API endpoints.
//4. to test the owner module, send a POST request to http://localhost:8000/owner/classes with the following JSON body:
/* {
  "name": "Yoga Class",
  "start_date": "2025-04-01T10:00:00Z",
  "end_date": "2025-04-01T11:00:00Z",
  "capacity": 20
}
*/
// for member to get the classes, send a GET request to http://localhost:8000/member/classes.
// This will return a list of classes available for booking.
// to book a class, send a POST request to http://localhost:8000/member/bookings with the following JSON body:
/* {
  "name": "Yoga Class",
  "date": "2025-12-01T10:00:00Z",
  "class_id": 2
}
// */
// This will book the class for the member and return a success message.
// If the class is not available or the booking details are invalid, it will return an error message.
// giving invalid class id or date outside the class date range will return an error message for member requests
