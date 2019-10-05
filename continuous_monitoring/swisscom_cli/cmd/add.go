package cmd

import (
	//"bufio"
	//"bytes"
	//"encoding/json"
	"fmt"
	"io/ioutil"

	//"net/http"
	//"os"
	//"reflect"

	"github.com/spf13/cobra"

	. "github.com/konyshev/swisscom_cli/models"
	. "github.com/konyshev/swisscom_cli/utils"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new setup",
	Long:  `Add a new setup`,
	Run: func(cmd *cobra.Command, args []string) {
		var inst Instance

		FillStruct(&inst)

		req, err := BuildRequest(&inst, config.Url, "POST")
		if err != nil {
			fmt.Println("Error during building request: ", err)
			return
		}

		resp, err := SendRequest(req)
		if err != nil {
			fmt.Println("Error during sending request: ", err)
			return
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error during parsing response: ", err)
			return
		}
		fmt.Println("response Body:", string(body))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
