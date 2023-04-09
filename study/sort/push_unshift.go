// 堆栈排序算法
package sort

func PushUnshiftSort(arr []int) []int {

	n := len(arr)
	newArr := make([]int, 0)

	for i := 0; i < n; i++ {
		len := len(newArr)
		if len == 0 {
			newArr = append(newArr, arr[i])
		} else {
			if arr[i] <= newArr[0] {
				newArr = append([]int{arr[i]}, newArr...)
			} else {
				newArr = append(newArr, arr[i])
			}
		}
	}

	return newArr
}
