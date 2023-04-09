package sort

func QuickSort(arr []int) []int {

	var newArr []int
	n := len(arr)
	if n <= 1 {
		return arr
	}

	var left, right []int
	for i := 1; i < n; i++ {
		if arr[i] < arr[0] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)

	for _, v := range left {
		newArr = append(newArr, v)
	}
	newArr = append(newArr, arr[0])
	for _, v := range right {
		newArr = append(newArr, v)
	}

	return newArr
}
