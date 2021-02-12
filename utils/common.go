package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

//GetJSON - ...
func GetJSON(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	FailOnError(err, "Failed to create new request")
	res, err := http.DefaultClient.Do(req)
	FailOnError(err, "Failed to get response")
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	FailOnError(err, "Failed to read response body")
	return body
}

//FailOnError - ...
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("Error: %s - Message: %s", err, msg)
	}
}

//SetDefaultValue - ...
func SetDefaultValue(src *string, dst string) {
	*src = dst
}
