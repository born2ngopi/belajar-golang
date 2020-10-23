package main

import (
	"fmt"
	//	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

func main() {
	img, err := imgio.Open("test.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	resized := transform.Resize(img, 512, 512, transform.Linear)

	if err := imgio.Save("output-512x512.png", resized, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
}
