package sorting

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	var low, high []int
	for _, i := range arr[1:] {
		if i < pivot {
			low = append(low, i)
			continue
		}
		high = append(high, i)
	}
	low = append(QuickSort(low), pivot)
	high = QuickSort(high)
	arr = append(low, high...)
	return arr
}
