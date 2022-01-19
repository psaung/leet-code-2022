package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := append(nums1, nums2...)
	sort.Ints(arr)
	len := len(arr)

	if len % 2 != 0 {
		return float64(arr[len/2])
	}

	return float64(arr[(len - 1) / 2] + arr[len / 2]) / 2;

}