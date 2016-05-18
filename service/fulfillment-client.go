package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type fulfillmentClient interface {
	getFulfillmentStatus(sku string) (status fulfillmentStatus, err error)
}

type fulfillmentWebClient struct {
	rootURL string
}

func (client fulfillmentWebClient) getFulfillmentStatus(sku string) (status fulfillmentStatus, err error) {
	httpclient := &http.Client{}

	req, _ := http.NewRequest("GET", client.rootURL+"/"+sku, nil)

	resp, err := httpclient.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	payload, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(payload, &status)
	if err != nil {
		fmt.Println("Failed to unmarshal server response.")
		return
	}

	return status, err
}
