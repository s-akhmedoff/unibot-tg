package stocks

import (
	"encoding/json"
	"fmt"

	"github.com/s-akhmedoff/unibot-tg/config"
	"github.com/s-akhmedoff/unibot-tg/utils"
)

type globalQuote struct {
	Symbol  string `json:"01. symbol"`
	Open    string `json:"02. open"`
	High    string `json:"03. high"`
	Low     string `json:"04. low"`
	Price   string `json:"05. price"`
	Volume  string `json:"06. volume"`
	PrevCl  string `json:"08. previous close"`
	Change  string `json:"09. change"`
	Percent string `json:"10. change percent"`
}

type stock struct {
	GlobalQuote globalQuote `json:"Global Quote,omitempty"`
}

func (s stock) String() string {
	if s.GlobalQuote.Symbol == "" {
		return "Failed to get stock, try another one"
	}

	result := fmt.Sprintf("Symbol: %s\nOpen: %s\nHigh: %s\tLow: %s\nPrice:"+
		" %s\nVolume: %s\nPrevious Close: %s\nChange: %s\tChange Percent: %s\n", s.GlobalQuote.Symbol,
		s.GlobalQuote.Open, s.GlobalQuote.High,
		s.GlobalQuote.Low, s.GlobalQuote.Price,
		s.GlobalQuote.Volume, s.GlobalQuote.PrevCl,
		s.GlobalQuote.Change, s.GlobalQuote.Percent)

	return result
}

// GetStocks - ...
func GetStocks(symbol string, config config.Config) string {
	if symbol == "" {
		utils.SetDefaultValue(&symbol, utils.StocksDefaultValue)
	}

	url := fmt.Sprintf(utils.StocksAPIURL, symbol, config.Currency.APIKey)

	stocks := new(stock)
	err := json.Unmarshal(utils.GetJSON(url), stocks)
	utils.FailOnError(err, "Failed to unmarshal data")

	if err != nil {
		return "Service is busy now!"
	}

	return stocks.String()
}
