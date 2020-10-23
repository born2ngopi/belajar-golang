package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func main() {

	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	// open the uploaded file
	file, err := os.Open("./img.png")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buff := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	_, err = file.Read(buff)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filetype := http.DetectContentType(buff)

	fmt.Println(filetype)

	switch filetype {
	case "image/jpeg", "image/jpg":
		fmt.Println(filetype)

	case "image/gif":
		fmt.Println(filetype)

	case "image/png":
		fmt.Println(filetype)

	case "application/pdf": // not image, but application !
		fmt.Println(filetype)
	default:
		fmt.Println("unknown file type uploaded")
	}

}
