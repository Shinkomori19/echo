package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

// struct data which will be returned to client.
// It' convention to use capital letter in struct.
// Correlating it to json by using `json:"key"`
type Data struct {
	Content string `json:"content"`
	Name string `json:"name"`
	Age int `json:"age"`
	Company string `json:"company"`
}

func main() {
	  // correlate path and handler function
    http.HandleFunc("/echo/post", postHandler)
		http.HandleFunc("/echo/get", getHandler)

		fmt.Print("Listening...")
		// .ListenAndServe() returns err only when it failed
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Decode json to object of Golang.
	var data Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "ERROR in decoding request json: %v", err)
	}
	// If we want to manipulate json sent from client, use decoded json here.

	json.NewEncoder(w).Encode(data)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// send default data
	shin := Data {
		Content: "responding from get endpoint",
		Name:  "Shin Komori",
		Age:    21,
		Company: "Soracom",
	}
	json.NewEncoder(w).Encode(shin)
}

// Previous version of main()
// func main() {
// 	// correlate path and handler function
// 	http.HandleFunc("/echo", echoHandler)

// 	fmt.Print("Listening...")
// 	// .ListenAndServe() returns err only when it failed
// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }

// func echoHandler(w http.ResponseWriter, r *http.Request) {
// // fmt.Fprint(w, r.Method)
// if err := r.ParseForm(); err != nil {
// 	fmt.Fprintf(w, "ERROR in .ParseForm(): %v", err)
// 	return
// }
// fmt.Fprintf(w, "\nReceived: %s\n", r.Form.Get("content"))
// }
