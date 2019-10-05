package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"

	. "github.com/konyshev/swisscom_cli/models"
)

func SendRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return resp, err
}

func BuildRequest(inst *Instance, url string, reqType string) (*http.Request, error) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(inst)
	req, err := http.NewRequest(reqType, url, buf)
	req.Header.Set("X-Custom-Header", "swisscom")
	req.Header.Set("Content-Type", "application/json")

	return req, err
}

func PrettyPrint(body []byte) {
	output := new(bytes.Buffer)
	json.Indent(output, body, "", "  ")
	fmt.Println(output)
}

func FillStruct(structToFill interface{}) {

	reflectStruct := reflect.ValueOf(structToFill).Elem()
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < reflectStruct.NumField(); i++ {
		field := reflectStruct.Field(i)
		fieldName := reflectStruct.Type().Field(i).Name

		if fieldName == "ID" {
			continue
		}
		if fieldName == "Contact" {
			var curContact Contact
			FillStruct(&curContact)
			field.Set(reflect.ValueOf(curContact))
			continue
		}

		fmt.Print("Enter value for " + fieldName + ":")
		value, _ := reader.ReadString('\n')

		field.Set(reflect.ValueOf(value))
	}
}
