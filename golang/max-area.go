package main


func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxArea(height []int) int {
	result, start, end := 0, 0, len(height) - 1;

	for start < end {
		shortest := min(height[start], height[end]) 

		result = max(result, shortest * (end - start))

		if(height[start] < height[end]) {
			start++;
		} else {
			end--;
		}
	}
	return result
}