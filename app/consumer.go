package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type News struct {
	ID     float32 `json:"id"`
	Title  string  `json:"title"`
	URL    string  `json:"url"`
	Source string  `json:"source"`
	Date   string  `json:"published_date"`
	Medias []Media `json:"media"`
}

type Media struct {
	MediaMetadatas []MediaMetadata `json:"media-metadata"`
}

type MediaMetadata struct {
	URL    string `json:"url"`
	Height string `json:"height"`
	Weight string `json:"weight"`
}

type NewsResponse struct {
	Status     string `json:"status"`
	Copyright  string `json:"copyright"`
	NumResults string `json:"num_results"`
	Results    []News `json:"results"`
}

func getNewYorkTimes() NewsResponse {
	response, err := http.Get("http://localhost:3000/ny")

	if err != nil {
		fmt.Println("Error")
	}

	var results NewsResponse
	_ = json.NewDecoder(response.Body).Decode(&results)

	// fmt.Println(results.Results[0].Medias)

	return results
}
