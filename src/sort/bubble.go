package sort

import "fmt"

func BubbleSort(arr *[]int) *[]int {
	arrVal := *arr

	for j := 0; j < len(*arr); j++ {
		for i := 0; i < len(*arr)-1-j; i++ {
			if arrVal[i] > arrVal[i+1] {
				temp := arrVal[i]
				arrVal[i] = arrVal[i+1]
				arrVal[i+1] = temp
			}
		}
	}

	return arr
}

func RunBubbleSortExample() {
	examples := []int{3, 1, 8, 4, 2, 6, 7}
	fmt.Println("Bubble sort examples:", BubbleSort(&examples))
}
