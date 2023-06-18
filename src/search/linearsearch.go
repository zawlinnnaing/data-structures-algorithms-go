package search

func LinearSearch(hayStack []int, needle int) bool {
	for i := 0; i < len(hayStack); i++ {
		if hayStack[i] == needle {
			return true
		}
	}
	return false
}
