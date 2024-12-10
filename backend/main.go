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

const apiKey = "" // Replace with your NewsAPI key
const newsAPIURL = "https://newsapi.org/v2/top-headlines?country=us&apiKey="

type NewsResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
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
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		lang = "en" // Default to English
	}

	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		http.Error(w, "API key not set", http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("%s%s&language=%s", newsAPIURL, apiKey, lang)

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news.Articles)
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
