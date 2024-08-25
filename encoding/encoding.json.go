package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"os"
)

func Unzip(name string) []byte {

	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	//Read the file
	gz, err := gzip.NewReader(file)
	if err != nil {
		log.Printf("file couldn't be read")
	}

	//copy the content of the files
	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	if err != nil {
		panic(err)
	}
	defer func(gz *gzip.Reader) {
		err := gz.Close()
		if err != nil {
		}
	}(gz)
	return buf.Bytes()

}
