package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func getApi(url string) ([]map[string]interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var geoData map[string]interface{}
	err = json.Unmarshal(data, &geoData)
	if err != nil {
		return nil, err
	}

	// Assuming `features` is an array of features as per standard GeoJSON format
	features, ok := geoData["features"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected structure, 'features' is not an array")
	}

	var results []map[string]interface{}
	for _, feature := range features {
		featureMap, ok := feature.(map[string]interface{})
		if !ok {
			continue // or handle error
		}
		results = append(results, featureMap)
	}

	return results, nil
}

func main() {
	const baseUrl = "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_week.geojson"

	results, err := getApi(baseUrl)
	if err != nil {
		log.Fatalf("Failed to get API data: %v", err)
	}

	// Inside your main function or wherever you're iterating over the results:
	for _, feature := range results {
		properties, ok := feature["properties"].(map[string]interface{})
		if !ok {
			continue // or handle error
		}

		place, ok := properties["place"].(string)
		if !ok {
			continue // or handle error
		}

		if strings.Contains(place, "California") {
			fmt.Println(place)
		}
	}
}
