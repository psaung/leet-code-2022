package main

func maxArea(height []int) int {
	result, len := height[0], len(height)

	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			min := height[i]
			if min > height[j]  {
				min = height[j]
			}
			if result < min * (j - 1) {
				result = min * (j - 1)
			}
		}
	}

	return result
}