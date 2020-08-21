package main

import (
	"./api/users"
	"encoding/json"
	"log"
	"net/http"
)

func handleReq() {
	http.HandleFunc("/users/staff", func(writer http.ResponseWriter, request *http.Request) {
		encoder := json.NewEncoder(writer)
		encoder.SetIndent("", "    ")
		encoder.Encode(users.GetEmployees())
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleReq()
}
