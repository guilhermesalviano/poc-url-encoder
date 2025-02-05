package application

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

func StartServer() {
    router := mux.NewRouter()

    corsMiddleware := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:5173", "http://172.0.0.1:5173"}),
        handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type"}),
        handlers.AllowCredentials(),
    )

    router.HandleFunc("/health", healthCheck).Methods("GET")
    router.HandleFunc("/encode", encodeHandler).Methods("POST")
    router.HandleFunc("/decode", decodeHandler).Methods("POST")

    http.ListenAndServe(":3333", corsMiddleware(router))
}