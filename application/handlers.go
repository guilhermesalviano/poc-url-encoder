package application

import (
    "encoding/json"
    "net/http"
)

// User data structure
type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users = []User{
    {ID: "1", Name: "John Doe", Email: "john@example.com"},
    {ID: "2", Name: "Jane Smith", Email: "jane@example.com"},
}

// Health check handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "API is running!"})
}

// Users handler
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}