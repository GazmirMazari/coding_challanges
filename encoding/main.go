package main

import (
	"github.com/bytedance/sonic"
	"log"
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

//
//func retrieveData(name string) {
//	file := Unzip(name)
//	var canRoot CanadaRoot
//	//unmarshall the data using st library
//	err := json.Unmarshal(file, &canRoot)
//	if err != nil {
//		log.Println("canadaRoot was not unmarshalled", err)
//	}
//
//}

func retrieveDatUsingByteDanceSonic(name string) {
	file := Unzip(name)
	var canRoot CanadaRoot
	//unmarshall the data using st library
	err := sonic.Unmarshal(file, &canRoot)
	if err != nil {
		log.Println("canadaRoot was not unmarshalled", err)
	}

}

func main() {
	const fileName = "canada_geometry.json.gz"
	//retrieveData(fileName)
	retrieveDatUsingByteDanceSonic(fileName)
}
