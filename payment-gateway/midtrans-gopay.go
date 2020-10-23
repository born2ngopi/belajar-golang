package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GopaySpec struct {
	PaymentType       string            `json:"payment_type"`
	TransactionDetail TransactionDetail `json:"transaction_details"`
	CustomExpiry      CustomExpiry      `json:"custom_expiry"`
	ItemDetails       []ItemDetails     `json:"item_details"`
}

type TransactionDetail struct {
	OrderId     string `json:"order_id"`
	GrossAmount int64  `json:"gross_amount"`
}

type CustomExpiry struct {
	Ordertime      string `json:"order_time"`
	ExpiryDuration int64  `json:"expiry_duration"`
	Unit           string `json:"unit"`
}

type ItemDetails struct {
	Id       string `json:"id"`
	Price    int64  `json:"price"`
	Quantity int64  `json:"quantity"`
	Name     string `json:"name"`
}

/// gopay response
type Gopay struct {
	StatusCode            string   `json:"status_code"`
	StatusMessage         string   `json:"status_message"`
	TransactionId         string   `json:"transaction_id"`
	OrderId               string   `json:"order_id"`
	GrossAmount           string   `json:"gross_amount"`
	PaymentType           string   `json:"payment_type"`
	TransactionTime       string   `json:"transaction_time"`
	TransactionStatus     string   `json:"transaction_status"`
	Actions               []Action `json:"actions"`
	ChanelResponseCode    string   `json:"chanel_response_code"`
	ChanelResponseMessage string   `json:"chanel_response_message"`
	Currency              string   `json:"currency"`
}

type Action struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Url    string `json:"url"`
}

func main() {

	var body = GopaySpec{
		PaymentType: "gopay",
		TransactionDetail: TransactionDetail{
			OrderId:     "6961",
			GrossAmount: 25000,
		},
		CustomExpiry: CustomExpiry{
			Ordertime:      time.Now().Format("2006-01-02 15:04:05 -0700"),
			ExpiryDuration: 15,
			Unit:           "minute",
		},
		ItemDetails: []ItemDetails{
			ItemDetails{
				Id:       "1",
				Price:    20000,
				Quantity: 1,
				Name:     "bakso",
			},
			ItemDetails{
				Id:       "2",
				Price:    2500,
				Quantity: 2,
				Name:     "krupuk",
			},
		},
	}

	requestBody, _ := json.Marshal(body)

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	uri := "https://api.sandbox.midtrans.com/v2/charge"
	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(requestBody))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic exampletokenuwu")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()

	fmt.Println(res.StatusCode)

	var resBody Gopay

	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resBody)
}
