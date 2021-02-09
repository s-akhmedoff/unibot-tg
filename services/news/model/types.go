package model

//Article - ...
type Article struct {
	Source        source `json:"source"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	URL           string `json:"url"`
	PublishedTime string `json:"publishedAt"`
}

type source struct {
	Name string `json:"name"`
}
