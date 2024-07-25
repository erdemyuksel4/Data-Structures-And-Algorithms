package sorting

func CountingSort(arr []int) []int {
	max := 0
	for _, a := range arr {
		if max < a {
			max = a
		}
	}
	countingArr := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		countingArr[arr[i]]++
	}
	var sortingArr []int
	for i, a := range countingArr {
		for j := 0; j < a; j++ {
			sortingArr = append(sortingArr, i)
		}
	}
	return sortingArr
}
