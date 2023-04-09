// 选择排序
package sort

func SelectionSort(a []int) []int {

	n := len(a)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}

		if min != i {
			temp := a[i]
			a[i] = a[min]
			a[min] = temp
		}
	}

	return a
}
