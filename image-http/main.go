package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/test", WebHandler)
	http.ListenAndServe(":8008", nil)
}

func WebHandler(w http.ResponseWriter, r *http.Request) {

	img, err := os.Open("example.jpg")
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	io.Copy(w, img)
}
