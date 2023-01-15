package utils

func DeleteItem[T any](arr []T, idx int) (r []T) {
	p := make([]T, idx)
	copy(p, arr[0:idx])
	end := make([]T, len(arr)-idx-1)
	copy(end, arr[idx+1:])
	return append(p, end...)
}

func InsertItem[T any](arr []T, idx int, item T) (r []T) {
	p := make([]T, idx)
	copy(p, arr[0:idx])
	end := make([]T, len(arr)-idx-1)
	copy(end, arr[idx:])

	r = append(p, item)
	r = append(r, end...)
	return
}
