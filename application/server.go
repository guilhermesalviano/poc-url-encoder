package application

import (
    "net/http"
    "github.com/gorilla/mux"
)

// StartServer initializes and runs the HTTP server
func StartServer() {
    router := mux.NewRouter()
    
    // Define routes
    router.HandleFunc("/", homeHandler).Methods("GET")
    router.HandleFunc("/users", getUsersHandler).Methods("GET")

    // Start server
    http.ListenAndServe(":3333", router)
}