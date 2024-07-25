package sorting

func BubbleSort(arrInt []int) []int {
	for j := 0; j < len(arrInt)-1; j++ {
		for i := 0; i < len(arrInt)-j-1; i++ {
			if arrInt[i] > arrInt[i+1] {
				arrInt[i], arrInt[i+1] = arrInt[i+1], arrInt[i]
			}
		}

	}
	return arrInt
}
