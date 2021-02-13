package definition

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"unibot-tg/config"
	"unibot-tg/utils"
)

type definition struct {
	Type       string `json:"model"`
	Definition string `json:"definition"`
	Example    string `json:"example,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
}

type definite struct {
	Definition    []definition `json:"definitions,omitempty"`
	Word          string       `json:"word"`
	Pronunciation string       `json:"pronunciation"`
}

func (d definite) String() string {
	var result string
	for i := range d.Definition {
		result += "Type: " + d.Definition[i].Type + "\nDefinition: " + d.Definition[i].Definition + "\n"
		if d.Definition[i].Example != "" {
			result += "Example: " + d.Definition[i].Example + "\n"
		}
		if d.Definition[i].ImageURL != "" {
			result += "Image URL: " + d.Definition[i].ImageURL
		}
		result += "\n"
	}

	return "Word: '" + d.Word + "' | Pronunciation: [" + d.Pronunciation + "]\n" + result
}

func getJSON(url, apiKey string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	utils.FailOnError(err, "Failed to create new request")
	req.Header.Add("Authorization", fmt.Sprintf(" Token %s", apiKey))
	res, err := http.DefaultClient.Do(req)
	utils.FailOnError(err, "Failed to get response")
	defer func() {
		err := res.Body.Close()
		utils.FailOnError(err, "Failed to close")
	}()
	body, err := ioutil.ReadAll(res.Body)
	utils.FailOnError(err, "Failed to read body")
	return body

}

//GetDefinition - Get definition of the word
func GetDefinition(arg string, config config.Config) string {
	if len(strings.Split(arg, " ")) > 1 {
		return "Too many words!\nUsage: /definition word"
	}
	if arg == "" {
		utils.SetDefaultValue(&arg, utils.DefinitionDefaultValue)
	}
	url := fmt.Sprintf(utils.DefinitionAPIURL, arg)

	definitionResult := new(definite)
	err := json.Unmarshal(getJSON(url, config.DefinitionKey), definitionResult)
	if err != nil {
		return "No definition was found for word: " + arg
	}
	utils.FailOnError(err, "Failed to unmarshal data")

	return definitionResult.String()
}
