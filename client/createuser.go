package main

import (
	"bufio"
	"fmt"
	"os"

	shoe "github.com/shoelick/goserver_example"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var newUserCmd = &cobra.Command{
	Use:   "useradd <username>",
	Short: "Create a new user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Password: ")
		passwd, _ := reader.ReadString('\n')

		c := shoe.NewShoeClient(viper.GetString("host"), []byte{}, true)
		err := c.CreateUser(args[0], passwd)

		return err
	},
}
