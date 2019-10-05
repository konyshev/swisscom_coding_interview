package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"

	. "github.com/konyshev/swisscom_cli/models"
	. "github.com/konyshev/swisscom_cli/utils"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Modify existing instance",
	Long:  `Modify existing instance`,
	Run: func(cmd *cobra.Command, args []string) {
		var reqStr = ""
		inst := &Instance{}

		if len(args) != 0 {
			reqStr = config.Url + "/" + args[0]
		} else {
			fmt.Println("Please specify an instance for update")
			return
		}

		resp, err := http.Get(reqStr)
		if err != nil {
			fmt.Println("Error during sending GET request: ", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		json.Unmarshal([]byte(body), inst)
		fmt.Print("Current value of Mail : " + inst.Email)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter new value for Mail :")
		value, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error during reading input value: ", err)
			return
		}

		fieldEmail := reflect.ValueOf(inst).Elem().FieldByName("Email")
		fieldEmail.Set(reflect.ValueOf(value))

		req, err := BuildRequest(inst, config.Url, "PUT")
		if err != nil {
			fmt.Println("Error during building request: ", err)
			return
		}

		respUpdate, err := SendRequest(req)
		if err != nil {
			fmt.Println("Error during sending PUT request: ", err)
			return
		}
		defer respUpdate.Body.Close()

		fmt.Println("response Status:", respUpdate.Status)
		bodyUpdate, _ := ioutil.ReadAll(respUpdate.Body)
		fmt.Println("response Body:", string(bodyUpdate))
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
