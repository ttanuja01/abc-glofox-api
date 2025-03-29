
---

# GloFox Management API

This repository contains an API with two main modules:

- **Owner Module**: Allows studio owners to create and manage classes.
- **Member Module**: Allows members to view and book classes.

The Project uses the **Gorilla Mux** router for handling HTTP requests and responses.

---

## Project Structure

The project has two core modules:

- **Owner Module**: For studio owners to create classes.
- **Member Module**: For members to view available classes and book them.

## Setup

### Prerequisites

Ensure that you have **Go** installed on your system. If not, you can download it from the official site: [Go Downloads](https://golang.org/dl/).

### Steps to Run the Application

1. **Clone this repository**:

   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Install Dependencies**:

   In your project folder, run the following command to install any missing dependencies:

   ```bash
   go mod tidy
   ```

3. **Run the Server**:

   Start the Go server with the following command:

   ```bash
   go run main.go
   ```

   This will start a local server on `http://localhost:8000`.

---

## Testing the API

You can use tools like **Thunder Client** or **Postman** to interact with the API.

### 1. **Owner Module: Create Classes**

To create a class as a **studio owner**, send a **POST** request to `http://localhost:8000/owner/classes` with the following **JSON body**:

```json
{
  "name": "Yoga Class",
  "start_date": "2025-04-01T10:00:00Z",
  "end_date": "2025-04-01T11:00:00Z",
  "capacity": 20
}
```

**Response**: The created class will be returned with the given details and a unique class id.

### 2. **Member Module: View Available Classes**

To view the available classes for booking, send a **GET** request to `http://localhost:8000/member/classes`.

**Response**: A list of all available classes for booking will be returned.

### 3. **Member Module: Book a Class**

To book a class as a **member**, send a **POST** request to `http://localhost:8000/member/bookings` with the following **JSON body**:

```json
{
  "name": "Ragnar Lothbrok",
  "date": "2025-04-01T10:00:00Z",
  "class_id": 6
}
```

**Response**: A success message will be returned confirming the booking.

If the class is not available or the booking details are invalid, an error message will be returned.

---

## Error Handling

- If an **invalid `class_id`** is provided, the API will return an error indicating that the class does not exist.
- If the **date is outside the class date range** (e.g., trying to book a class before or after its start and end dates), the API will return an error.
- Any other **invalid data** (such as missing fields or malformed JSON) will result in a `400 Bad Request` response with a descriptive error message.

---

## Notes

- The **Owner Module** does not have any authentication in the current version.
- The **Member Module** assumes the member provides valid class IDs and booking dates.

---