// 插入排序
package sort

func InsertionSort(arr []int) []int {

	n := len(arr)

	for i := 1; i < n; i++ {
		get := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > get {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = get
	}

	return arr
}
