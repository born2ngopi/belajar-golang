package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hallo", hallo)
	fmt.Println("running")
	http.ListenAndServe("127.0.0.1:3030", nil)

}

func hallo(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Message string `json:"message"`
	}{
		"hallo gaes",
	}

	json.NewEncoder(w).Encode(resp)
}
