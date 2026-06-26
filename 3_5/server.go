package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]string{"message": "hello, world!"})
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		panic(err)
		return
	}
}
