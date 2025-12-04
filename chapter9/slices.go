package main

import (
	"fmt"
)

// slice: dynamic array,use it when size is unknown
func main() {
	// uninitialised slice is nil
	// var nums []int
	// fmt.Println(nums==nil)
	// fmt.Println(len(nums))

	// fmt.Println(cap(nums))
	// fmt.Println(nums)

	// nums = append(nums, 2)
	// nums = append(nums, 12)
	// nums = append(nums, 23)
	// nums = append(nums, 21)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// var nums = make([]int, 0, 4)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:2])
	// fmt.Println(nums[:2])
	// fmt.Println(nums[1:])

	// slice package

	// var nums1 = []int{1, 2}
	// var nums2 = []int{1, 2,3}

	// fmt.Println(slices.Equal(nums1, nums2))

	var nums = [][]int{{1, 4, 2}, {2, 3, 5}}
	fmt.Println(nums)
}
