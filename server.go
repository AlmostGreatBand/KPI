package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var port = ":8795"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		type Json struct {
			Time string `json:"time"`
		}
		jsn := Json{ Time: time.Now().Format(time.RFC3339) }
		if er := json.NewEncoder(w).Encode(jsn); er != nil {
			fmt.Fprintf(w, "an error appeared: %v", er)
		}
	})

	fmt.Printf("Server on localhost%v", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
