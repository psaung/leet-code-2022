package main

import "strings"

func romanToInt(s string) int {
	t := strings.Split(s, "") 
	 m, len, result := make(map[string]int), len(t), 0 
	m["I"] = 1	
	m["V"] = 5 
	m["X"] = 10 
	m["L"] = 50 
	m["C"] = 100 
	m["D"] = 500 
	m["M"] = 1000 

	for i := 0; i < len; i++ {
		if( i < len - 1 && m[t[i]] < m[t[i + 1]]) {
			result += m[t[i+1]] - m[t[i]]
			i += 1;
		} else {
			result += m[t[i]]
		}

	}

	return result
}