package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)


type Response struct {
	Code		int	`json:"code"`
	Message		string	`json:"message"`
}



func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(Response{ 404, "Not Found"})

		fmt.Println(r)
	})

	http.ListenAndServe(":8080", mux)
}
