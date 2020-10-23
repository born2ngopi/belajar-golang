package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	//"io"
	"image/png"
	"net/http"
)

func main() {

	http.HandleFunc("/", test)

	fmt.Println("vim-go")
	http.ListenAndServe(":8000", nil)
}

func test(w http.ResponseWriter, r *http.Request) {

	q, err := qrcode.New(`{"order_id":"sdjasduew329eiw9ud98sdhsjdhuia", "venue_id":"jdiadahsdd82381283182hds","date":"dasdjs8888sjdisdj8"}`, qrcode.Medium)
	if err != nil {
		fmt.Println("qrcode error")
		return
	}

	img := q.Image(256)

	w.Header().Set("Content-Type", "image/png")
	//	io.Copy(w, png)
	png.Encode(w, img)
}
