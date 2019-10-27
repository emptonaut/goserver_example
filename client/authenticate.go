package main

import (
	"bufio"
	"fmt"
	"os"

	shoe "github.com/shoelick/goserver_example"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authCmd = &cobra.Command{
	Use:   "auth <username>",
	Short: "Create new session and token with the server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Password: ")
		passwd, _ := reader.ReadString('\n')

		c := shoe.NewShoeClient(viper.GetString("host"), []byte{}, true)
		tok, err := c.Authenticate(args[0], passwd)

		if err == nil {
			fmt.Println(tok)
		}

		return err
	},
}
