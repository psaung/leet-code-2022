package main

import "fmt"

func letterCombinations(digits string) []string {
	result, phones := []string{}, map[string]string {
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}; 

	if(len(digits) < 1) {
		return result;
	}

	var combination func (string, map[string]string, string, int)

	combination = func(digits string, phones map[string]string, prefix string, depth int) {
		if depth == len(digits) {
			result = append(result, prefix)
			fmt.Println(result)
			return
		}

		for i := 0; i < len(phones[string(digits[depth])]); i++ {
			combination(digits, phones, prefix + string(phones[string(digits[depth])][i]), depth+1)
		}
	}

	combination(digits, phones, "", 0)
	return result
}