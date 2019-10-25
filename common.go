package goserver

// RequestData is our general purpose container for data sent for this server.
// Different field may be used for different endpoints.
type RequestData struct {
	Token    string
	Username string
	Password string
}
