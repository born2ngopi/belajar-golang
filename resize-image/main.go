package main

import (
	"fmt"
	"html/template"
	//	"io"
	"mime/multipart"
	"net/http"
	"os"
	//	"path/filepath"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		var tmpl = template.Must(template.ParseFiles("index.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/process", uploadImage)

	http.ListenAndServe(":4000", nil)
}

func validate(f multipart.File) {

	buff := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	_, err := f.Read(buff)

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

func uploadImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//alias := r.FormValue("alias")

	m := r.MultipartForm

	file := m.File["file"]
	fmt.Println(len(file))
	for i, _ := range file {

		f, err := file[i].Open()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		validate(f)

	}

	//	uploadedFile, handler, err := r.FormFile("file")
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//	defer uploadedFile.Close()

	//	dir, err := os.Getwd()
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}

	//	filename := handler.Filename
	//	if alias != "" {
	//		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	//	}

	//	fileLocation := filepath.Join(dir, "files", filename)
	//	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//	defer targetFile.Close()

	//	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}

	w.Write([]byte("done"))
}
