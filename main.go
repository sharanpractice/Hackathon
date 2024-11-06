package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
)

// GreetHandler handles requests to the /greet endpoint
func GreetHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    response := map[string]string{"message": "Hello, welcome to the Go Greet API!"}
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// GetPort returns the port for the server to listen on, defaulting to 8080 if not set
func GetPort() string {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    return port
}

func main() {
    port := GetPort()
    
    http.HandleFunc("/greet", GreetHandler)
    
    log.Printf("Starting server on port %s...", port)
    if err := http.ListenAndServe(":" + port, nil); err != nil {
        log.Fatalf("Could not start server: %v\n", err)
    }
}
