package main

type Result struct {
	ArticleID      string   `json:"article_id"`
	Title          string   `json:"title"`
	Link           string   `json:"link"`
	Keywords       []string `json:"keywords"`
	Creator        []string `json:"creator"`
	VideoURL       any      `json:"video_url"`
	Description    string   `json:"description"`
	Content        string   `json:"content"`
	PubDate        string   `json:"pubDate"`
	PubDateTZ      string   `json:"pubDateTZ"`
	ImageURL       any      `json:"image_url"`
	SourceID       string   `json:"source_id"`
	SourcePriority int      `json:"source_priority"`
	SourceName     string   `json:"source_name"`
	SourceURL      string   `json:"source_url"`
	SourceIcon     string   `json:"source_icon"`
	Language       string   `json:"language"`
	Country        []string `json:"country"`
	Category       []string `json:"category"`
	AiTag          string   `json:"ai_tag"`
	Sentiment      string   `json:"sentiment"`
	SentimentStats string   `json:"sentiment_stats"`
	AiRegion       string   `json:"ai_region"`
	AiOrg          string   `json:"ai_org"`
	Duplicate      bool     `json:"duplicate"`
}
type NewsResponse struct {
	Status       string   `json:"status"`
	TotalResults int      `json:"totalResults"`
	Results      []Result `json:"results"`
	NextPage     string   `json:"nextPage"`
}
