package utils

func FindIndex[T comparable](arr []T, el T) int {
	for idx, element := range arr {
		if element == el {
			return idx
		}
	}
	return -1
}
