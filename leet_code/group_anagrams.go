package leet_code

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)

	for _, word := range strs {
		var count [26]int
		for _, c := range word {
			count[c-'a']++
		}
		anagramMap[count] = append(anagramMap[count], word)
	}

	result := make([][]string, len(anagramMap))
	idx := 0
	for _, v := range anagramMap {
		result[idx] = v
		idx++
	}
	return result
}
