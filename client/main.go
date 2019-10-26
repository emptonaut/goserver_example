package main

import (
	"fmt"

	shoe "github.com/shoelick/goserver_example"
)

func main() {

	c := shoe.NewShoeClient("localhost", []byte{}, true)
	//out, err := c.RequestSecret()
	resp, err := c.Request("/secret", &shoe.RequestData{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}

}
