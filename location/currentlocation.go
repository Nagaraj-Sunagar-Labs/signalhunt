package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationResponse struct {
	Location struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lng"`
	} `json:"location"`
	Accuracy float64 `json:"accuracy"`
}

func main() {
	apiKey := "GMAP_API_KEY" // Replace with your actual API key

	requestBody := []byte(`{"considerIp": true}`)

	url := "https://www.googleapis.com/geolocation/v1/geolocate?key=" + apiKey
	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var location LocationResponse
	if err := json.NewDecoder(response.Body).Decode(&location); err != nil {
		panic(err)
	}

	fmt.Println("Latitude : ", location.Location.Latitude)
	fmt.Println("Longitude: ", location.Location.Longitude)
	fmt.Println("Accuracy: ", location.Accuracy, "meters")

	mapsURL := fmt.Sprintf(
		"https://www.google.com/maps?q=%.6f,%.6f",
		location.Location.Latitude,
		location.Location.Longitude,
	)
	fmt.Println("Open this URL:", mapsURL)
}
