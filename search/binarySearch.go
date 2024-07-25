package search

func BinarySearch(arr []int, target int) int {
	len := len(arr)
	mid := len/2 + len%2
	if arr[mid] < target {
		return BinarySearch(arr[:mid], target)
	} else if arr[mid] > target {
		return BinarySearch(arr[mid:], target)
	}
	return mid
}
