package main

import (
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar!")
}

func main() {
	fmt.Println("vim-go")
	var env = os.Getenv("PORT")

	var port = fmt.Sprintf(":%s", env)

	http.HandleFunc("/", index)

	http.ListenAndServe(port, nil)
}
