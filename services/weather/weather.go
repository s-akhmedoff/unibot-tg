package weather

import (
	"encoding/json"
	"fmt"
	"time"
	"unibot-tg/config"
	"unibot-tg/utils"
)

type weather struct {
	Sys struct {
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunset"`
	} `json:"sys"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Name string      `json:"name"`
	Cod  interface{} `json:"cod"`
}

func (w weather) String() string {
	switch w.Cod {
	case "404":
		return "Incorrect City!"
	case 200:
	}
	return fmt.Sprintf("\tWeather"+
		"\nHey there! Now in %s '%s', I mean %s"+
		"\nTemperature is %s °C, but feels like %s °C"+
		"\nPressure: %s hPa\tHumidity: %s %%"+
		"\nAlso Wind Speed is: %s mPs"+
		"\nYou should know that:"+
		"\nSunrise at: %s"+
		"\nSunset at: %s", w.Name, w.Weather[0].Main, w.Weather[0].Description, fmt.Sprintf("%.2f", w.Main.Temp),
		fmt.Sprintf("%.2f", w.Main.FeelsLike), fmt.Sprintf("%d", w.Main.Pressure), fmt.Sprintf("%d", w.Main.Humidity),
		fmt.Sprintf("%.2f", w.Wind.Speed), time.Unix(int64(w.Sys.Sunrise), 0).Format("01.02 15:04"),
		time.Unix(int64(w.Sys.Sunset), 0).Format("01.02 15:04"))
}

//GetWeather - Get weather in `country` provided by OpenWeatherMap
func GetWeather(country string, config config.Config) string {
	if country == "" {
		utils.SetDefaultValue(&country, utils.WeatherDefaultValue)
	}
	url := fmt.Sprintf(utils.WeatherAPIURL, country, config.WeatherKey)

	ByteData := utils.GetJSON(url)

	WeatherObj := new(weather)
	ParseError := json.Unmarshal(ByteData, &WeatherObj)
	utils.FailOnError(ParseError, "Failed to parse")

	return WeatherObj.String()
}
