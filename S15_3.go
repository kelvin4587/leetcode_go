package main

import "sort"

func threeSum(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	answer := make([][]int, 0)
	for first := 0; first < length; first++ {
		if first == 0 || nums[first] != nums[first-1] {
			third := length - 1
			for second := first + 1; second < length; second++ {
				if second == first+1 || nums[second] != nums[second-1] {
					for second < third && nums[first]+nums[second]+nums[third] > 0 {
						third--
					}
					if second == third {
						break
					}
					if nums[first]+nums[second]+nums[third] == 0 {
						answer = append(answer, []int{nums[first], nums[second], nums[third]})
					}
				}
			}
		}
	}
	return answer
}
