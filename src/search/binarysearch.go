package search

import "math"

func BinarySearch(arr []int, value int) bool {
	low, high := 0, len(arr)-1

	for low <= high {
		middlePoint := low + ((high - low) / 2)
		middleValue := arr[middlePoint]
		if middleValue == value {
			return true
		}
		if value > middleValue {
			low = middlePoint + 1
		} else {
			high = middlePoint - 1
		}
	}

	return false
}

func TwoCrystalBallsProblem(results []bool) int {
	jumpAmount := int(math.Sqrt(float64(len(results))))
	i := 0
	for ; i < len(results); i += jumpAmount {
		result := results[i]
		if result {
			break
		}
	}

	for j := i - jumpAmount; j < i && j < len(results)-1; j++ {
		if results[j] {
			return j
		}
	}

	return -1
}
