package algorythmns

func binarySearch(list []int, target int) int {

	low := 0
	high := len(list) - 1
	mid := (low + high) / 2

	guess := list[mid]

	for {
		if guess == target {
			return guess
		}
		if guess < target {
			low = mid + 1
		}

		if guess > target {
			high = mid + 1
		}

		return 0

	}

}
