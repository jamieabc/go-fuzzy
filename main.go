package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"./protocol"
)

const (
	scheme  string = "https"
	host    string = "127.0.0.1:2131"
	rpcPath string = "/bitmarkd/rpc"
)

var rpcURL = url.URL{
	Scheme: scheme,
	Host:   host,
	Path:   rpcPath,
}

func generateJSON(params protocol.DataGeneration) (io.Reader, error) {
	params.GenRandomData()
	params.JustifyData()

	// keep for test
	// params.SampleData()

	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	paramsStr := string(paramsJSON)
	fmt.Println("post body: ", paramsStr)
	return strings.NewReader(paramsStr), nil
}

// send http request
func sendRequest(reqBody io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", rpcURL.String(), reqBody)
	if err != nil {
		fmt.Printf("error message: %s\n", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("error message: %s\n", err)
		return nil, err
	}
	return resp, nil
}

func main() {
	// choose format
	// params := &protocol.ProvenanceRpc{}
	// params := protocol.New(protocol.TransactionStatusType)
	// params := protocol.New(protocol.NodeInfoType)
	params := protocol.New(protocol.BitmarksProofType)

	for true {
		body, err := generateJSON(params)

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
		fmt.Printf("%s\n\n", string(response))

		// wait a bit
		time.Sleep(50 * time.Millisecond)
	}
}
