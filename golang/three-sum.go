package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	length := len(nums)
    left, right := 0, 0
	
    var triplets [][]int

	for i := 0; i < length - 2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			left = i + 1
			right = length - 1;
			for left < right {
				sum := nums[i] + nums[left] + nums[right];
				if sum == 0 {

					aaa := []int{nums[i], nums[left], nums[right]}
					triplets = append(triplets, aaa)
					for left + 1 < length && nums[left] == nums[left + 1] {
						left++
					}

					for right - 1 < 0 && nums[right] == nums[right - 1] {
						right--
					}

					left++
					right--
				} else if sum < 0 {
					left++
				} else if sum > 0 {
					right --
				}
			}
		}
	}

	return triplets
}