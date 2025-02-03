package application

import (
    "net/http"
    "github.com/gorilla/mux"
)

func StartServer() {
    router := mux.NewRouter()

    router.HandleFunc("/health", healthCheck).Methods("GET")
    router.HandleFunc("/encode", encodeHandler).Methods("POST")
    router.HandleFunc("/decode", decodeHandler).Methods("GET")

    http.ListenAndServe(":3333", router)
}