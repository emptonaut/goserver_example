package goserver

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// RequestData is our general purpose container for data sent between client and server.
// Different field may be used for different endpoints.
type RequestData struct {
	Token    string
	Username string
	Password string
	Error    string
}

// ParseRequestData is used to unmarshal data in requests and responses.
func ParseRequestData(body io.ReadCloser) (*RequestData, error) {
	reqBody, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	data := &RequestData{}
	err = json.Unmarshal(reqBody, data)
	if err != nil {
		return nil, err
	}
	return data, err
}
