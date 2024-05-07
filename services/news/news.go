package news

import (
	"encoding/json"
	"fmt"

	"github.com/s-akhmedoff/unibot-tg/config"
	"github.com/s-akhmedoff/unibot-tg/services/news/model"
	"github.com/s-akhmedoff/unibot-tg/utils"
)

type news struct {
	Articles []model.Article `json:"articles"`
}

func (n news) String() string {
	var titles string

	if len(n.Articles) < 1 {
		return "Incorrect source, please try another source i.e 'yandex.ru' or 'usatoday.com'"
	}

	for i := range n.Articles {
		pbTime := n.Articles[i].PublishedTime[11:19]
		titles += pbTime + " -> " + n.Articles[i].Title + "\n"
	}

	return titles
}

// GetNews - ...
func GetNews(domain string, config config.Config) string {
	utils.SetDefaultValue(&domain, utils.NewsDefaultValue)

	url := fmt.Sprintf(utils.NewsAPIURL, domain, config.News.APIKey)

	newsCollection := new(news)
	err := json.Unmarshal(utils.GetJSON(url), newsCollection)
	utils.FailOnError(err, "Failed to unmarshal data")

	return newsCollection.String()
}
