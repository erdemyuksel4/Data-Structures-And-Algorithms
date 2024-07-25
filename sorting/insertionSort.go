package sorting

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			for j, a := range arr[:i] {
				if arr[i] < a {
					arr = insert(arr, j, arr[i])
					break
				}
			}
			arr = pop(arr, i+1)
		}
	}
	return arr
}
func insert(arr []int, index int, num int) []int {

	newArr := make([]int, len(arr)+1)
	copy(newArr[:index], arr[:index])

	newArr[index] = num
	copy(newArr[index+1:], arr[index:])

	return newArr
}
func pop(arr []int, index int) []int {
	var newArr []int
	newArr = arr[:index]
	newArr = append(newArr, arr[index+1:]...)
	return newArr
}
