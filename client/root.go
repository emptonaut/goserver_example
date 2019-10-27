package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "shoeclient",
		Short: "A demonstration client for this example server",
	}
)

func init() {

	rootCmd.PersistentFlags().String("host", "localhost", "Hostname of the server")
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))

	rootCmd.AddCommand(authCmd, secretCmd, logoutCmd, passwdCmd, newUserCmd)

}

/*func main() {

	c := shoe.NewShoeClient(, []byte{}, true)
	//out, err := c.RequestSecret()
	//resp, err := c.Request("/secret", &shoe.RequestData{})
	//resp, err := c.CreateUser("michael", "dude")
	resp, err := c.Authenticate("michael", "dude")
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

}*/
