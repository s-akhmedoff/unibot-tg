package definition

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/s-akhmedoff/unibot-tg/config"
	"github.com/s-akhmedoff/unibot-tg/utils"
)

type definition struct {
	Type       string `json:"model"`
	Definition string `json:"definition"`
	Example    string `json:"example,omitempty"`
	ImageURL   string `json:"image_url,omitempty"`
}

type definite struct {
	Definition    []definition `json:"definitions,omitempty"`
	Word          string       `json:"word"`
	Pronunciation string       `json:"pronunciation"`
}

func (d definite) String() string {
	var result string

	for _, def := range d.Definition {
		result += "Type: " + def.Type + "\nDefinition: " + def.Definition + "\n"

		if def.Example != "" {
			result += "Example: " + def.Example + "\n"
		}

		if def.ImageURL != "" {
			result += "Image URL: " + def.ImageURL
		}

		result += "\n"
	}

	return "Word: '" + d.Word + "' | Pronunciation: [" + d.Pronunciation + "]\n" + result
}

func getJSON(url, apiKey string) []byte {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	utils.FailOnError(err, "Failed to create new request")

	req.Header.Add("Authorization", " Token "+apiKey)

	res, err := http.DefaultClient.Do(req)
	utils.FailOnError(err, "Failed to get response")

	defer func() {
		err := res.Body.Close()

		utils.FailOnError(err, "Failed to close")
	}()

	body, err := io.ReadAll(res.Body)
	utils.FailOnError(err, "Failed to read body")

	return body
}

// GetDefinition - Get definition of the word.
func GetDefinition(arg string, config config.Config) string {
	if len(strings.Split(arg, " ")) > 1 {
		return "Too many words!\nUsage: /definition word"
	}

	if arg == "" {
		utils.SetDefaultValue(&arg, utils.DefinitionDefaultValue)
	}

	url := fmt.Sprintf(utils.DefinitionAPIURL, arg)

	definitionResult := new(definite)

	err := json.Unmarshal(getJSON(url, config.Definition.APIKey), definitionResult)
	if err != nil {
		return "No definition was found for word: " + arg
	}

	utils.FailOnError(err, "Failed to unmarshal data")

	return definitionResult.String()
}
