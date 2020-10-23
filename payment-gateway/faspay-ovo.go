package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type FaspaySpec struct {
	Request        string     `json:"request"` // charge venue
	MerchantId     string     `json:"merchant_id"`
	Merchant       string     `json:"merchant"`        // dolin
	BillNo         string     `json:"bill_no"`         // order number
	BillDate       string     `json:"bill_date"`       // order date now
	BillExpired    string     `json:"bill_expired"`    // order data + 15 minute
	BillDesc       string     `json:"bill_desc"`       // charge venue
	BillCurrency   string     `json:"bill_currency"`   // IDR
	BillTotal      string     `json:"bill_total"`      // gross_amount
	PaymentChannel string     `json:"payment_channel"` //
	PayType        string     `json:"pay_type"`        // 1
	CustNo         string     `json:"cust_no"`         // user id
	CustName       string     `json:"cust_name"`       // user name
	Msisdn         string     `json:"msisdn"`          // nomor telephone
	Email          string     `json:"email"`           // user email
	Terminal       string     `json:"terminal"`        // 21
	Items          []ItemSpec `json:"item"`
	Reserve1       string     `json:"reserve1"`
	Reserve2       string     `json:"reserve2"`
	Signature      string     `json:"signature"`
}

type ItemSpec struct {
	Id      string `json:"id"`
	Product string `json:"product"`
	Qty     string `json:"qty"`
	Amount  string `json:"amount"`
	Type    string `json:"type"` // ticket
}

type OvoSpec struct {
	Response     string     `json:"response"`
	TrxId        string     `json:"trx_id"`
	MerchantId   string     `json:"merchant_id"`
	Merchant     string     `json:"merchant"`
	BillNo       string     `json:"bill_no"`
	BillItem     []BillItem `json:"bill_item"`
	ResponseCode string     `json:"response_code"`
	ResponseDesc string     `json:"response_desc"`
	Deeplink     string     `json:"deeplink"`
	WebUrl       string     `json:"web_url"`
	RedirectUrl  string     `json:"redirect_url"`
}

type BillItem struct {
	Id          string `json:"id"`
	Product     string `json:"product"`
	Qty         string `json:"qry"`
	Amount      string `json:"amount"`
	PaymentPlan string `json:"payment_plan"`
	MarchanId   string `json:"merchan_id"`
	Tenor       string `json:"tenor"`
	Type        string `json:"type"`
	Uri         string `json:"uri"`
	ImageUrl    string `json:"image_url"`
}

const (
	MERCHANT_ID       string = "merchant_id"
	MERCHANT_PASSWORD string = "password"
)

func main() {

	now := time.Now()

	text := fmt.Sprintf("bot%s%s%s", MERCHANT_ID, MERCHANT_PASSWORD, "123456789")
	hashMd5 := fmt.Sprintf("%x", md5.Sum([]byte(text)))
	sha := sha1.New()
	sha.Write([]byte(hashMd5))
	encrypted := sha.Sum(nil)
	signature := fmt.Sprintf("%x", encrypted)

	var body = FaspaySpec{
		Request:        "Post Data Transaksi",
		MerchantId:     MERCHANT_ID,
		Merchant:       "Dolin",
		BillNo:         "123456789",
		BillDate:       now.Format("2006-01-02 15:04:05"),
		BillExpired:    now.Local().Add(time.Minute * time.Duration(15)).Format("2006-01-02 15:04:05"),
		BillDesc:       "pembelian barang",
		BillCurrency:   "IDR",
		BillTotal:      "50000",
		PaymentChannel: "812",
		PayType:        "1",
		CustNo:         "2",
		CustName:       "chandra",
		Msisdn:         "",
		Email:          "",
		Terminal:       "21",
		Items: []ItemSpec{
			ItemSpec{
				Id:      "001",
				Product: "product",
				Qty:     "1",
				Amount:  "",
				Type:    "tiket",
			},
		},
		Reserve1:  "",
		Reserve2:  "",
		Signature: signature,
	}

	requestBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("error 1")
		fmt.Println(err.Error())
		return
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	uri := "https://dev.faspay.co.id/cvr/300011/10"
	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error 2")
		fmt.Println(err.Error)
		return
	}
	defer res.Body.Close()

	fmt.Println(res.StatusCode)

	var resBody OvoSpec

	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		fmt.Println("error ovo")
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resBody)
	fmt.Printf("\n\n")

	uri = "https://dev.faspay.co.id/pws/ovo_direct"
	dreq, _ := http.NewRequest(http.MethodPost, uri, nil)
	dreq.Header.Set("Content-Type", "application/json")
	dreq.Header.Set("trx_id", resBody.TrxId)
	dreq.Header.Set("ovo_number", "085669610513")
	dreq.Header.Set("signature", signature)

	dres, err := client.Do(dreq)
	if err != nil {
		fmt.Println("error 3")
		fmt.Println(err.Error())
		return
	}
	defer dres.Body.Close()

	fmt.Println(dres.StatusCode)

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bodyString := string(bodyByte)
	fmt.Println(bodyString)

	//var dresBody map[string]interface{}

	//err = json.NewDecoder(dres.Body).Decode(&dresBody)
	//if err != nil {
	//	fmt.Println("error 4")
	//	log.Fatalln(err.Error())
	//	return
	//}

	//fmt.Println(dresBody)
}
