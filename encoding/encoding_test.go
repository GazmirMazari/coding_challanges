package main

import (
	bytes2 "bytes"
	"testing"
)

//func BenchmarkAlgorithmOne(b *testing.B) {
//	var buff bytes2.Buffer
//	nameFile := "canada_geometry.json.gz"
//	b.ResetTimer()
//	for n := 0; n < b.N; n++ {
//		buff.Reset()
//		retrieveData(nameFile)
//	}
//}

func BenchmarkAlgorithmTwo(b *testing.B) {
	var buff bytes2.Buffer
	nameFile := "canada_geometry.json.gz"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buff.Reset()
		retrieveDatUsingByteDanceSonic(nameFile)
	}
}
