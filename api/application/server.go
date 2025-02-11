package application

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"

    "os"
)

func StartServer() {
    router := mux.NewRouter()

    webHost := os.Getenv("WEB_HOST")
    if webHost == "" {
        webHost = "http://localhost:5173"
    }

    corsMiddleware := handlers.CORS(
        handlers.AllowedOrigins([]string{webHost}),
        handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type"}),
        handlers.AllowCredentials(),
    )

    router.HandleFunc("/health", healthCheck).Methods("GET")
    router.HandleFunc("/encode", encodeHandler).Methods("POST")
    router.HandleFunc("/decode", decodeHandler).Methods("POST")

    port := os.Getenv("SERVER_PORT")
    if port == "" {
        port = "3333"
    }

    http.ListenAndServe(":"+port, corsMiddleware(router))
}