package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	shoe "github.com/shoelick/goserver_example"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var passwdCmd = &cobra.Command{
	Use:   "passwd <username> <token>",
	Short: "Change user password",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("New Password: ")
		passwd, _ := reader.ReadString('\n')

		c := shoe.NewShoeClient(viper.GetString("host"), []byte{}, true)
		err := c.ChangePasswd(args[0], passwd, args[1])

		if err == nil {
			log.Println("OK")
		}

		return err
	},
}
