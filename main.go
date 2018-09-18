package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"local/random"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type dataGeneration interface {
	justifyData(methodName string)
	genRandomData()
	sampleData()
}

type rpcFormat struct {
	ID     string             `json:"id,omitempty"`
	Method string             `json:"method,omitempty"`
	Params []provenanceFormat `json:"params,omitempty"`
}

type provenanceFormat struct {
	Count uint   `json:"count,omitempty"`
	TxID  string `json:"txId,omitempty"`
}

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

func (rpc *rpcFormat) justifyData(methodName string) {
	rpc.Method = methodName
	rpc.ID = "1"
	for _, p := range rpc.Params {
		p.Count = 20
	}
}

// genRandomData generates random data fits specific interface
func (rpc *rpcFormat) genRandomData() {
	r := random.New()
	r.Fuzz(rpc)
}

func generateJSON(params dataGeneration, methodName string) (io.Reader, error) {
	params.genRandomData()
	params.justifyData(methodName)

	// keep for test
	// params.sampleData()

	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	paramsStr := string(paramsJSON)
	fmt.Println("post body: ", paramsStr)
	return strings.NewReader(paramsStr), nil
}

// sampleData generates correct data
func (rpc *rpcFormat) sampleData() {
	rpc.ID = "1"
	rpc.Method = "Bitmark.Provenance"
	rpc.Params = []provenanceFormat{
		provenanceFormat{
			Count: 20,
			TxID:  "2dc8770718b01f0205ad991bfb4c052f02677cff60e65d596e890cb6ed82c861",
		},
	}
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
	params := &rpcFormat{}

	for true {
		body, err := generateJSON(params, "Bitmark.Provenance")

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
