package sorting

func SelectionSort(arr []int) []int {
	index := 0
	len := len(arr)

	for k := 0; k < len-1; k++ {
		index = k
		for i := k + 1; i < len; i++ {
			if arr[i] < arr[index] {
				index = i
			}
		}
		arr[k], arr[index] = arr[index], arr[k]
	}

	return arr
}
