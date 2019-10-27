package main

import (
	"fmt"
	"io/ioutil"

	shoe "github.com/shoelick/goserver_example"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var secretCmd = &cobra.Command{
	Use:   "secret <token>",
	Short: "Get secret stuff available only to those who have logged in",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		c := shoe.NewShoeClient(viper.GetString("host"), []byte{}, true)
		resp, err := c.Request("/secret", &shoe.RequestData{
			Token: args[0],
		})

		if err != nil {
			return err
		}

		if resp.Status != "200 OK" {
			log.Errorf("Got status %s", resp.Status)
		}

		if err != nil {
			log.Println(err)
		} else {
			if respBody, err := ioutil.ReadAll(resp.Body); err == nil {
				fmt.Println(string(respBody))
			} else {
				log.Error(err)
			}
		}

		return err
	},
}
