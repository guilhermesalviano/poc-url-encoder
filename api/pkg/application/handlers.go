package application

import (
    "encoding/json"
    "net/http"
    "net/url"
)

type RequestBody struct {
    Content string `json:"content"`
}

func healthCheckOld(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "API is running!"})
}

func encodeHandlerOld(w http.ResponseWriter, r *http.Request) {
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

    if body.Content == "" {
        http.Error(w, `{"error": "Content parameter is required"}`, http.StatusBadRequest)
        return
    }

    parsedContent, err := url.Parse(body.Content)
    if err != nil && encodeOnlyParams {
        http.Error(w, `{"error": "Invalid URL format"}`, http.StatusBadRequest)
        return
    }

    var encodedURL string

    switch {
    case encodeOnlyParams:
        query := parsedContent.Query()
        encodedQuery := make(url.Values)
        
        for key, values := range query {
            // encodedKey := url.QueryEscape(key)
            encodedKey := key
            for _, value := range values {
                // encodedQuery.Add(encodedKey, url.QueryEscape(value))
                encodedQuery.Add(encodedKey, value)
            }
        }

        parsedContent.RawQuery = encodedQuery.Encode()
        encodedURL = parsedContent.String()

    default:
        encodedURL = url.QueryEscape(body.Content)
    }

    response := map[string]interface{}{
        "original_content":    body.Content,
        "content":     encodedURL,
        "parameters_used": map[string]bool{
            "encode_only_params": encodeOnlyParams,
        },
    }

    json.NewEncoder(w).Encode(response)
}

func decodeHandlerOld(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method != http.MethodPost {
        http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
        return
    }

    var body RequestBody
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
        return
    }

    contentDecode, err := url.QueryUnescape(body.Content)
	if err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
        return
	}

    response := map[string]interface{}{
        "original_content":    body.Content,
        "content":     contentDecode,
    }

    json.NewEncoder(w).Encode(response)
}