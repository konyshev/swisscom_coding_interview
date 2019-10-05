package cmd

import (
	//"bytes"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/konyshev/swisscom_cli/utils"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Show instances",
	Long:  `Show all instances. To see a specific instance please use "swisscom_cli ls "<instance_id>"`,
	Run: func(cmd *cobra.Command, args []string) {
		var reqStr = ""
		if len(args) != 0 {
			reqStr = config.Url + "/" + args[0]
		} else {
			reqStr = config.Url
		}

		resp, err := http.Get(reqStr)
		if err != nil {
			fmt.Println("Error during sending request: ", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error during parsing response: ", err)
			return
		}

		PrettyPrint(body)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
