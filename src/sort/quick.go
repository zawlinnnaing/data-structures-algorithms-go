package sort

func partition(arr *[]int, low int, high int) int {
	pivotIndex := ((high - low) / 2) + low
	pivot := (*arr)[pivotIndex]
	left := low
	right := high

	for left <= right {
		for (*arr)[left] < pivot {
			left++
		}
		for (*arr)[right] > pivot {
			right--
		}

		if left > right {
			break
		}

		temp := (*arr)[left]
		(*arr)[left] = (*arr)[right]
		(*arr)[right] = temp

		left++
		right--
	}

	return pivotIndex
}

func quickS(arr *[]int, low int, high int) {
	if low >= high {
		return
	}

	pivotIndex := partition(arr, low, high)
	quickS(arr, low, pivotIndex-1)
	quickS(arr, pivotIndex+1, high)
}

func Quick(input []int) {
	quickS(&input, 0, len(input)-1)
}
