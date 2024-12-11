package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const newsDATAURL = "https://newsdata.io/api/1/latest?"

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

// CORS Middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func fetchNewsHandler(w http.ResponseWriter, r *http.Request) {
	// lang := r.URL.Query().Get("lang")
	// if lang == "" {
	// 	lang = "en" // Default to English
	// }

	// country := r.URL.Query().Get("cat")
	// if country == "" {
	// 	country = "us" // Default to USA
	// }

	apiKey := os.Getenv("NEWS_DATA_API_KEY")
	if apiKey == "" {
		http.Error(w, "API key not set", http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("%sapikey=%s&country=us&size=1", newsDATAURL, apiKey)
	fmt.Println(url)

	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch news", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var news NewsResponse
	err = json.NewDecoder(response.Body).Decode(&news)
	if err != nil {
		http.Error(w, "Failed to parse news data", http.StatusInternalServerError)
		return
	}
	fmt.Println(w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news.Results)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "pong"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/ping", pingHandler)
	r.HandleFunc("/news", fetchNewsHandler)

	port := "5432"
	fmt.Printf("Server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, enableCORS(r)))
}
