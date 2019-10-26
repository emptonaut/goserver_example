package main

import (
	"fmt"
	"io/ioutil"

	shoe "github.com/shoelick/goserver_example"
)

func main() {

	c := shoe.NewShoeClient("localhost", []byte{}, true)
	//out, err := c.RequestSecret()
	resp, err := c.Request("/secret", &shoe.RequestData{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Response:")
		fmt.Println(resp)
		fmt.Println("Body:")

		respBody, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(respBody))
		fmt.Println(err)
	}

}
