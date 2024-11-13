package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string `json:"id"`
	OriginalURL  string `json:"origin_url"`
	ShortURL     string `json:"short_url"`
	CreationDate string `json:"created_date"`
}

/*
	d9736711 -->{
	    ID: "d9736711",
	    OriginURL: "https://github.com/CodeByAnkita",
	    ShortURL: "9736711",
	    CreationDate: "time.Now()",

}
*/
var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL)) // It converts the originalURL string to a byte slice
	fmt.Println("hasher: ", hasher)

	data := hasher.Sum(nil)
	fmt.Println("hasher data: ", data)
	hash := hex.EncodeToString(data)
	fmt.Println("hash:", hash)
	fmt.Println("hash final:", hash[:8])

	return hash[:8]
}
func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	id := shortURL // Use the short URL as the ID for simplicity
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	return shortURL
}
func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	shortURL := createURL(data.URL)
	// fmt.Fprintf(w, shortURL) 
	response := struct {
		shortURL string `json:"shortURL"`
	}{shortURL: shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	fmt.Println("Starting URL shortener...")
	OriginalURL := "https://github.com/CodeByAnkita/"
	generateShortURL(OriginalURL)

	// func main() {
	// Register the handler function to handle all requests to the root URL ("/")
	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", ShortURLHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
