package application

import (
    "encoding/json"
    "net/http"
    "os"
    "log"

    "url-encoder/pkg/application/rabbitmq"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

// URL struct for encoding/decoding requests
type URL struct {
    OriginalURL string `json:"originalUrl"`
    EncodedURL  string `json:"encodedUrl,omitempty"`
}

var rmq *rabbitmq.RabbitMQ

func init() {
    uri := os.Getenv("RABBITMQ_URI")
    if uri == "" {
        uri = "amqp://guest:guest@localhost:5672/"
    }

    var err error
    rmq, err = rabbitmq.NewRabbitMQ(uri)
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }

    // Create queue for URL encoding
    queueName := "url_encoder_queue"
    err = rmq.CreateQueue(queueName)
    if err != nil {
        log.Fatalf("Failed to create queue: %v", err)
    }
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func encodeHandler(w http.ResponseWriter, r *http.Request) {
    var urlRequest URL
    err := json.NewDecoder(r.Body).Decode(&urlRequest)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Publish URL to RabbitMQ for encoding
    queueName := "url_encoder_queue"
    err = rmq.Publish(queueName, []byte(urlRequest.OriginalURL))
    if err != nil {
        http.Error(w, "Failed to queue URL for encoding", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusAccepted)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "URL queued for encoding test",
        "status": "processing",
    })
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
    var urlRequest URL
    err := json.NewDecoder(r.Body).Decode(&urlRequest)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // In a real-world scenario, you'd have a separate service 
    // or process to handle decoding and return the result
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Decoding not implemented in this example",
    })
}

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

    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(router)))
}


// Graceful shutdown function (optional but recommended)
func Shutdown() {
    if rmq != nil {
        rmq.Close()
    }
}