package main

import (
	"fmt"
)

func swap(nums []int, i int, j int) []int {
	nums[i] ^= nums[j]
	nums[j] ^= nums[i]
	nums[i] ^= nums[j]

	return nums
}

func maxHeapify(nums []int, i int) []int {
	left := (i * 2) + 1
	right := (i * 2) + 2
	largest := i

	if left < len(nums) && nums[left] > nums[largest] {
		largest = left
	}

	if right < len(nums) && nums[right] > nums[largest] {
		largest = right
	}

	if largest != i {
		nums = swap(nums, largest, i)
		nums = maxHeapify(nums, largest)
	}

	return nums
}

func buildMaxHeap(nums []int) []int {
	for i := len(nums) / 2; i >= 0; i-- {
		nums = maxHeapify(nums, i)
	}

	return nums
}

func heapSort(nums []int) []int {
	heapSize := len(nums) - 1
	nums = buildMaxHeap(nums)
	for i := heapSize; i >= 1; i-- {
		swap(nums, 0, i)
		buildMaxHeap(nums[0:i])
	}

	return nums
}

func main() {
	fmt.Println(heapSort([]int{1, 4, 3, 5, 2, 6, 5, 7, 6, 8}))
	fmt.Println(heapSort([]int{6, 5, 4, 3, 2, 1}))
}
