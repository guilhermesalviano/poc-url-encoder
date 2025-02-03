package application

import (
    "encoding/json"
    "net/http"
    "net/url"
)

type RequestBody struct {
    URL string `json:"url"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "API is running!"})
}

func encodeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method != http.MethodPost {
        http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
        return
    }

    queryParams := r.URL.Query()
    encodeOnlyParams := queryParams.Get("encode_only_params") == "true"

    var body RequestBody
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
        return
    }

    if body.URL == "" {
        http.Error(w, `{"error": "URL parameter is required"}`, http.StatusBadRequest)
        return
    }

    parsedURL, err := url.Parse(body.URL)
    if err != nil {
        http.Error(w, `{"error": "Invalid URL format"}`, http.StatusBadRequest)
        return
    }

    var encodedURL string

    switch {
    case encodeOnlyParams:
        query := parsedURL.Query()
        encodedQuery := make(url.Values)
        
        for key, values := range query {
            encodedKey := url.QueryEscape(key)
            for _, value := range values {
                encodedQuery.Add(encodedKey, url.QueryEscape(value))
            }
        }
        
        parsedURL.RawQuery = encodedQuery.Encode()
        encodedURL = parsedURL.String()

    default:
        encodedURL = url.QueryEscape(body.URL)
    }

    response := map[string]interface{}{
        "original_url":    body.URL,
        "encoded_url":     encodedURL,
        "parameters_used": map[string]bool{
            "encode_only_params": encodeOnlyParams,
        },
    }

    json.NewEncoder(w).Encode(response)
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	encodedURL := r.URL.Query().Get("url")
	if encodedURL == "" {
		http.Error(w, `{"error": "URL parameter is required"}`, http.StatusBadRequest)
		return
	}

	decodedURL, err := url.QueryUnescape(encodedURL)
	if err != nil {
		http.Error(w, `{"error": "Invalid encoded URL"}`, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"decoded_url": decodedURL,
	})
}