package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		type Json struct{
			Time string `json:"time"`
		}
		q:=time.Now()

		json.NewEncoder(w).Encode(q)
	})
	log.Fatal(http.ListenAndServe(":8795", nil))
}
