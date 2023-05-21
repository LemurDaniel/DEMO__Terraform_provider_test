package client

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	address string
	http    *http.Client
}

func Create(address string) *Client {
	return &Client{
		http:    &http.Client{},
		address: address,
	}
}

/*
(c *Client):
	allows to call this method on an instance of the Client struct.

client:
 	variable of an instance of type *Client (a pointer to Client).

Get():
	method with no args

(io.ReadCloser, error):
	return type of method. Pointer to key/value map, and error

*/

func (client *Client) Get(path string) (io.ReadCloser, error) {
	return client.request(path, "GET", bytes.Buffer{})
}

func (client *Client) request(endpoint, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {

	fullAdress := fmt.Sprintf("%s/%s", client.address, endpoint)
	request, err := http.NewRequest(method, fullAdress, &body)
	if err != nil {
		return nil, err
	}

	//TODO
	//req.Header.Add("Authorization", client.authToken)
	request.Header.Add("Content-Type", "application/json")

	response, err := client.http.Do(request)
	if err != nil {
		return nil, err
	}

	log.Default().Output(1, fmt.Sprintf("[CLIENT] -- HOST: %s", request.URL.Host))
	log.Default().Output(1, fmt.Sprintf("[CLIENT] -- PATH: %s", request.URL.Path))
	log.Default().Output(1, fmt.Sprintf("[CLIENT] -- RESPONSECODE: %s", response.Status))

	if response.StatusCode != http.StatusOK {
		//respBody := new(bytes.Buffer)
		//_, err := respBody.ReadFrom(resp.Body)

		return nil, fmt.Errorf("Statuscode: %v", response.StatusCode)
	}

	//responseBody, err := ioutil.ReadAll(response.Body)
	//log.Default().Output(1, fmt.Sprintf("[CLIENT] -- RESPONSECONTENT: %s", string(responseBody)))
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}
