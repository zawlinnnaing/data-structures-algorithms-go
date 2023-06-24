package util

type comparableSliceType interface {
	string | int | float32 | float64
}

func SliceEquals[T comparableSliceType](slice1 []T, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
