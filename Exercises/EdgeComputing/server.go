package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "time"
)

// handler for the /api/v1/events endpoint
func eventsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Read the delay from the environment variable
    delayStr := os.Getenv("REQUEST_DELAY_MS")
    delay, err := strconv.Atoi(delayStr)
    if err != nil {
        delay = 0 // Default delay to 0 if parsing fails
    }

    // Add the delay before processing the request
    time.Sleep(time.Duration(delay) * time.Millisecond)

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Event received"))
}

func main() {
    http.HandleFunc("/api/v1/events", eventsHandler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not specified
    }

    fmt.Printf("Server is running on port %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}