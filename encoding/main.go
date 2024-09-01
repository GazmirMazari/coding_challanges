package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/bytedance/sonic"
	"io"
	"log"
	"os"
)

type CanadaRoot struct {
	Type     string `json:"type"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			Name string `json:"name"`
		} `json:"properties"`
		Geometry struct {
			Type        string         `json:"type"`
			Coordinates [][][2]float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

func algoOne(name string) {
	var canRoot CanadaRoot

	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the gzipped file
	gz, err := gzip.NewReader(file)
	if err != nil {
		log.Printf("file couldn't be read: %v", err)
		return
	}
	defer gz.Close()

	// Copy the content of the file into a buffer
	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	if err != nil {
		panic(err)
	}

	// Unmarshal the data using the standard library
	err = json.Unmarshal(buf.Bytes(), &canRoot)
	if err != nil {
		log.Println("canadaRoot was not unmarshalled:", err)
		return
	}

	// Print or process `canRoot` as needed
	log.Printf("algoOne processed: %+v\n", canRoot)
}

func algoTwo(name string) {
	var canRoot CanadaRoot

	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the gzipped file
	gz, err := gzip.NewReader(file)
	if err != nil {
		log.Printf("file couldn't be read: %v", err)
		return
	}
	defer gz.Close()

	// Copy the content of the file into a buffer
	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	if err != nil {
		panic(err)
	}

	// Unmarshal the data using Sonic library
	err = sonic.Unmarshal(buf.Bytes(), &canRoot)
	if err != nil {
		log.Println("canadaRoot was not unmarshalled:", err)
		return
	}

	// Print or process `canRoot` as needed
	log.Printf("algoTwo processed: %+v\n", canRoot)
}

func main() {
	const fileName = "canada_geometry.json.gz"
	algoOne(fileName)
	algoTwo(fileName)
}
