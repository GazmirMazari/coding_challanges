package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// URL for querying California earthquakes
	url := "https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=2024-01-01&endtime=2024-08-24&minlatitude=32&maxlatitude=42&minlongitude=-125&maxlongitude=-114&minmagnitude=2.5"

	// Send GET request to the API
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to grab the URL: %v", err)
		return
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Unmarshal the JSON response into a map
	var earthquakeData map[string]interface{}
	if err := json.Unmarshal(body, &earthquakeData); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Access the "features" key, which should be an array of objects
	features, ok := earthquakeData["features"].([]interface{})
	if !ok {
		log.Fatal("Unexpected format for 'features'")
	}

	// Collect the earthquake codes
	var codes []string
	for _, feature := range features {
		featureMap, ok := feature.(map[string]interface{})
		if !ok {
			continue
		}

		// Access properties within each feature
		properties, ok := featureMap["properties"].(map[string]interface{})
		if ok {
			code, ok := properties["code"].(string)
			if ok {
				codes = append(codes, code)
			}
		}
	}

	// Print the collected codes
	fmt.Println("These are the earthquake codes:", codes)
}
