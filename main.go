package main

import (
	"fmt"
	"net/http"
)

// Plumbing: This function "carries data" like pipes carry water.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "🚀 Backend magic! Your request worked!") // What the user sees indirectly
}

// Electrical: The server "powers" the app.
func main() {
	fmt.Println("🔌 Starting the server at port 8080") // What the user sees directly
	http.HandleFunc("/", handler)                     // Connect the "pipe" (API endpoint)
	http.ListenAndServe(":8080", nil)                 // "Turn on" the server (port 8080), port is like a house number
}
