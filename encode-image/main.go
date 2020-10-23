package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("vim-go")

	f, err := os.Open("./example.jpg")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err.Error())
	}

	// encode as base64
	encoded := base64.StdEncoding.EncodeToString(content)
	err = ioutil.WriteFile("./encode-example.txt", []byte(encoded), 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
