package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/0x-Parzival/blockplain-voting/backend/router"  // Import router package
)

func main() {
    // Initialize the router
    r := mux.NewRouter()

    // Call SetupRouter from the router package
    router.SetupRouter(r)

    // Start the server
    fmt.Println("Starting server on port 8080...")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal("Error starting server: ", err)
    }
}
