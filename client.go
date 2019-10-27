package goserver

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// ShoeClient is a convenience wrapper around http.Client
type ShoeClient struct {
	client    *http.Client
	serverURI string // exclude protocol
}

// NewShoeClient configures a ShoeClient for communicating with a ShoeServer
func NewShoeClient(host string, CA []byte, skipVerify bool) *ShoeClient {

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipVerify},
	}

	return &ShoeClient{
		serverURI: host,
		client: &http.Client{
			Timeout:   time.Second * 10,
			Transport: transport,
		},
	}
}

func (c *ShoeClient) CreateUser(username, password string) (err error) {

	data := &RequestData{
		Username: username,
		Password: password,
	}

	var resp *http.Response
	if resp, err = c.Request("/user/create", data); err != nil {
		return errors.New(resp.Status)
	}

	if resp.Status == "200 OK" {

		data, err = ParseRequestData(resp.Body)
		if data.Error != "" {
			err = errors.New(data.Error)
		}
	}

	return
}

func (c *ShoeClient) Authenticate(username, password string) (string, error) {
	data := &RequestData{
		Username: username,
		Password: password,
	}

	var reqData *RequestData
	if resp, err := c.Request("/user/authenticate", data); err != nil {
		return "", err
	} else {
		if resp.Status != "200 OK" {
			return "", fmt.Errorf("Got status %s", resp.Status)
		} else if reqData, err = ParseRequestData(resp.Body); err != nil {
			return "", err
		}
	}
	return reqData.Token, nil
}

// RequestSecret requests secrets from a ShoeServer
func (c *ShoeClient) RequestSecret() (string, error) {

	data := &RequestData{}
	req, err := c.newRequest(data, "/secrets")
	if err != nil {
		return "", err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	var respBody []byte
	respBody, err = ioutil.ReadAll(resp.Body)
	return string(respBody), err
}

// endpoint _should_ include the leading slash
func (c *ShoeClient) newRequest(data *RequestData, endpoint string) (*http.Request, error) {

	jsonStr, err := json.Marshal(data)
	url := fmt.Sprintf("https://%s%s", c.serverURI, endpoint)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	return req, err
}

// RequestSecret requests secrets from a ShoeServer
func (c *ShoeClient) Request(endpoint string, data *RequestData) (*http.Response, error) {

	req, err := c.newRequest(data, endpoint)
	if err != nil {
		return nil, err
	}
	return c.client.Do(req)
}
