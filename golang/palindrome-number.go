package main

import "strconv"

func isPalindrome(x int) bool {
	t := []rune(strconv.Itoa(x))
	left, right :=  0, len(t) - 1

	for left < right {
		if t[left] != t[right] {
			return false
		}
		left++
		right--
	}

	return true;
}