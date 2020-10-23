package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	client := &http.Client{
		Timeout: 100 * time.Second,
	}

	var reqBody = struct {
		To       string `json:"to"`
		Username string `json:"username"`
		Password string `json:"password"`
		Content  string `json:"content"`
	}{
		"082364100456",
		"smppsuser",
		"smppsusr",
		"woy",
	}

	requestBody, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(http.MethodPost, "https://sms.dolin.id/send", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()

	fmt.Println(res.StatusCode)

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	bodyString := string(bodyByte)
	fmt.Println(bodyString)
}
