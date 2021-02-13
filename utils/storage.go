package utils

const (
	//WeatherAPIURL - ...
	WeatherAPIURL = "https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s"
	//WeatherDefaultValue - ...
	WeatherDefaultValue = "Tashkent"

	//NewsAPIURL - ...
	NewsAPIURL = "http://newsapi.org/v2/everything?domains=%s&pageSize=30&apiKey=%s"
	//NewsDefaultValue - ...
	NewsDefaultValue = "yandex.ru"

	//StocksAPIURL - ...
	StocksAPIURL = "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s"
	//StocksDefaultValue - ...
	StocksDefaultValue = "AMD"

	//DefinitionAPIURL - ...
	DefinitionAPIURL = "https://owlbot.info/api/v4/dictionary/%s"
	//DefinitionDefaultValue - ...
	DefinitionDefaultValue = "Word"

	//CurrencyAPIURL - ...
	CurrencyAPIURL = "https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=%s&to_currency=%s&apikey=%s"
	//CurrencyDefaultValueFrom - ...
	CurrencyDefaultValueFrom = "USD"
	//CurrencyDefaultValueTo - ...
	CurrencyDefaultValueTo = "UZS"

	//HelpDefaultResponse - ...
	HelpDefaultResponse = "Supported commands:\n" +
		"/weather city - Fetch weather in city, default: Tashkent\n" +
		"/news domain - Show latest news from domain, default yandex.ru\n" +
		"/stocks symbol - Observe symbol's stocks, default AMD\n" +
		"/currency from to - Currency from to, default USD -> UZS\n" +
		"/definition word - Look up definition of the word, default word\n" +
		"/cli command - Execute command on the host\n" +
		"/generator n - Generate secure n-length password, default 15"
)
