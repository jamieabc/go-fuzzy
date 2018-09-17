package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"local/random"
	"net/http"
	"strings"
	"time"
)

type rpcFormat struct {
	ID     int              `json:"id,omitempty"`
	Method string           `json:"method,omitempty"`
	Params provenanceFormat `json:"params,omitempty"`
}

type provenanceFormat struct {
	Count uint   `json:"count,omitempty"`
	TxID  string `json:"txId,omitempty"`
}

const (
	host string = "127.0.0.1"
	port string = "2131"
	path string = "/bitmarkd/rpc"
	url  string = "https://" + host + ":" + port + path
)

// genRandomData generates random data fits specific interface
func genRandomData(params interface{}, methodName string) (io.Reader, error) {
	fmt.Println("params: ", params)
	r := random.New()
	r.Fuzz(params)

	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	paramsStr := string(paramsJSON)
	return strings.NewReader(paramsStr), nil
}

// sampleData generates correct data
func sampleData() io.Reader {
	return strings.NewReader(`{"id":"1","method":"Bitmark.Provenance","params":[{"count":20,"txId":"2dc8770718b01f0205ad991bfb4c052f02677cff60e65d596e890cb6ed82c861"}]}`)
}

// send http request
func sendRequest(reqBody io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		fmt.Printf("error message: %s\n", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error message: %s\n", err)
		return nil, err
	}
	return resp, nil
}

func main() {
	// choose format
	params := &rpcFormat{}

	for true {
		body, err := genRandomData(params, "Bitmark.Provenance")

		if nil != err {
			fmt.Println("Error message: ", err)
			return
		}

		resp, err := sendRequest(body)
		if err != nil {
			fmt.Printf("error message: %s\n", err)
			return
		}
		defer resp.Body.Close()

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error message: %s\n", err)
			return
		}
		fmt.Printf("response: %s\n\n", response)

		// wait a bit
		time.Sleep(50 * time.Millisecond)
	}
}
