package main

import (
	"log"

	shoe "github.com/shoelick/goserver_example"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var logoutCmd = &cobra.Command{
	Use:   "logout <token>",
	Short: "Logout user and delete session",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := shoe.NewShoeClient(viper.GetString("host"), []byte{}, true)
		var err error
		if err = c.Logout(args[0]); err == nil {
			log.Println("Done.")
		}
		return err
	},
}
