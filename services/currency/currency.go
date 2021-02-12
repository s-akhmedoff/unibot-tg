package currency

import (
	"encoding/json"
	"fmt"
	"strings"
	"unibot-tg/config"
	"unibot-tg/utils"
)

type exchanger struct {
	RCE struct {
		FromC string `json:"1. From_Currency Code"`
		ToC   string `json:"3. To_Currency Code"`
		Ex    string `json:"5. Exchange Rate"`
		Bid   string `json:"8. Bid Price"`
		Ask   string `json:"9. Ask Price"`
	} `json:"Realtime Currency Exchange Rate"`
}

func (e exchanger) String() string {
	if e.RCE.FromC == "" {
		return "Error"
	}

	return e.RCE.Ex[:5] + " " + e.RCE.ToC
}

// GetCurrency - get currency rate
func GetCurrency(arg string, config config.Config) string {
	if arg == "" {
		utils.SetDefaultValue(&arg, fmt.Sprintf("%s %s", utils.CurrencyDefaultValueFrom, utils.CurrencyDefaultValueTo))
	}
	url := fmt.Sprintf(utils.CurrencyApiUrl, strings.Split(arg, " ")[0], strings.Split(arg, " ")[1], config.CurrencyKey)

	exchange := new(exchanger)
	err := json.Unmarshal(utils.GetJSON(url), &exchange)
	utils.FailOnError(err, "Failed to unmarshal")

	return exchange.String()
}
