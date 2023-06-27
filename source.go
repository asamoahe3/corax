package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var products DummyProducts

// fucn MakeGetRequest() makes http get requst to a define api endpoint
// @param{url string} accept the url path of the api endpoint
// @praram{reciever}: a struct type to decode the response.
func MakeGetRequest(url string, reciever interface{}) (err error) {
	res, err := http.Get(url)
	if err == nil {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, reciever)
	}

	return
}

func init() {

	//Get request to dummyjson.com/products endpoint
	if err := MakeGetRequest("https://dummyjson.com/products", &products); err != nil {
		log.Fatal(err)
	}
}
