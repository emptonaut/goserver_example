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

// TODO move comment to server context or readme
// The better solution to something like this would be to rewrite the http.Handler interface to accept
// a third argument for common data that every endpoint would require, such as session data.
// However, that's beyond the scope of this example exercise. For now, every endpoint must
// reparse the RequestData even if the root receive (ServerHTTP) already parsed it. Known design flaw.
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
