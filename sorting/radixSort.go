package sorting

func RadixSort(arr []int) []int {
	radixArr := make([][]int, 10)
	digitIndex := 0
	max := 0
	for _, a := range arr {
		if a > max {
			max = a
		}
	}
	maxLen := len(digits(max))

	for ; maxLen > digitIndex; digitIndex++ {
		for _, a := range arr {
			digit := 0
			if len(digits(a))-1 >= digitIndex {
				digit = digits(a)[digitIndex]
			}
			radixArr[digit] = append(radixArr[digit], a)
		}
		arr = []int{}
		for j := 0; j < 10; j++ {
			arr = append(arr, radixArr[j]...)
		}
		radixArr = make([][]int, 10)
	}

	return arr
}
func digits(n int) []int {
	var digits []int
	for n > 0 {
		digits = append(digits, int(n%10))
		n /= 10
	}
	return digits
}
