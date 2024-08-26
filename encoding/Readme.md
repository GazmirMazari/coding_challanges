Command run for comparison of these test is
go test -run none -bench AlgorithmOne -benchtime 3s -benchmem
go test -run none -bench AlgorithmTwo -benchtime 3s -benchmem

Benchmark result
BenchmarkAlgorithmTwo-8              999           4566118 ns/op         1760972 B/op       1182 allocs/op  -standard library/json
BenchmarkAlgorithmOne-8              364          10538601 ns/op         1443300 B/op       1675 allocs/op - sonic/bytedance


BenchmarkAlgorithmTwo: Faster algorithm (4.57 ms/op), uses more memory (1.76 MB/op), and has fewer memory allocations (1182 allocs/op).

BenchmarkAlgorithmOne: Slower algorithm (10.54 ms/op), uses less memory (1.44 MB/op), but has more memory allocations (1675 allocs/op).

#Create memory profiles to understand memory allocation
go test -run none -bench AlgorithmOne -benchtime 3s -benchmem -memprofile mem_algo_one.out
go test -run none -bench AlgorithTwo -benchtime 3s -benchmem -memprofile mem_algo_two.out

go tool pprof -alloc_space memcpu.test mem_algo_one.out
(pprof) list retrieveData
Total: 934.79MB
ROUTINE ======================== jsonEncoder.retrieveData 
0   931.28MB (flat, cum) 99.62% of Total
.          .     23:func retrieveData(name string) {
.   710.05MB     24:   file := Unzip(name)
.          .     25:   var canRoot CanadaRoot
.          .     26:   //unmarshall the data using st library
.   221.23MB     27:   err := json.Unmarshal(file, &canRoot)
.          .     28:   if err != nil {
.          .     29:           log.Println("canadaRoot was not unmarshalled", err)
.          .     30:   }
.          .     31:
.          .     32:}
ROUTINE ======================== 
0   931.28MB (flat, cum) 99.62% of Total
.          .      8:func BenchmarkAlgorithmOne(b *testing.B) {
.          .      9:   var buff bytes2.Buffer
.          .     10:   nameFile := "canada_geometry.json.gz"
.          .     11:   b.ResetTimer()
.          .     12:   for n := 0; n < b.N; n++ {
.          .     13:           buff.Reset()
.   931.28MB     14:           retrieveData(nameFile)
.          .     15:   }
.          .     16:}
.          .     17:
.          .     18:func BenchmarkAlgorithmTwo(b *testing.B) {
.          .     19:   var buff bytes2.Buffer

ROUTINE ======================== jsonEncoder.retrieveDatUsingByteDanceSonic in C:\Users\gazmi\repos\algos-challange\encoding\main.go
512.02kB     1.77GB (flat, cum) 99.83% of Total
.          .     34:func retrieveDatUsingByteDanceSonic(name string) {
.     1.07GB     35:   file := Unzip(name)
512.02kB   512.02kB     36:   var canRoot CanadaRoot
.          .     37:   //unmarshall the data using st library
.   711.08MB     38:   err := sonic.Unmarshal(file, &canRoot)
.          .     39:   if err != nil {
.          .     40:           log.Println("canadaRoot was not unmarshalled", err)
.          .     41:   }
.          .     42:
.          .     43:}


ROUTINE ======================== jsonEncoder.BenchmarkAlgorithmTwo in C:\Users\gazmi\repos\algos-challange\encoding\encoding_test.go
0     1.77GB (flat, cum) 99.83% of Total
.          .     18:func BenchmarkAlgorithmTwo(b *testing.B) {
.          .     19:   var buff bytes2.Buffer
.          .     20:   nameFile := "canada_geometry.json.gz"
.          .     21:   b.ResetTimer()
.          .     22:   for n := 0; n < b.N; n++ {
.          .     23:           buff.Reset()
.     1.77GB     24:           retrieveDatUsingByteDanceSonic(nameFile)
.          .     25:   }
.          .     26:}


go build -gcflags "-m -m"