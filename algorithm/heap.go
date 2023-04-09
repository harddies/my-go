package main

import "fmt"

func main() {
	nums := []int{5, 3, 4, 2, 6, 8, 1}
	fmt.Println(kthLargest(nums, 2))
	heapSort(nums)
}

func heapSort(nums []int) {
	buildMaxHeap(nums, len(nums))
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeap(nums, 0, i)
	}
	fmt.Println(nums)
}

func kthLargest(nums []int, k int) int {
	size := len(nums)
	buildMaxHeap(nums, size)
	fmt.Println(nums)

	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		size--
		maxHeap(nums, 0, size)
	}
	fmt.Println(nums)
	return nums[0]
}

func buildMaxHeap(nums []int, size int) {
	for i := size / 2; i >= 0; i-- {
		maxHeap(nums, i, size)
	}
}

func maxHeap(nums []int, i, size int) {
	left, right, largest := 2*i+1, 2*i+2, i
	if left < size && nums[left] > nums[largest] {
		largest = left
	}
	if right < size && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeap(nums, largest, size)
	}
}
