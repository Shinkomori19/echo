package main

import (
    "fmt"
    "net/http"
		"log"
)

func main() {
	  // correlate path and handler function
    http.HandleFunc("/echo", echoHandler)

		fmt.Print("Listening...")
		// .ListenAndServe() returns err only when it failed
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, r.Method)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ERROR in .ParseForm(): %v", err)
		return
	}
	fmt.Fprintf(w, "\nReceived: %s\n", r.Form.Get("content"))
}