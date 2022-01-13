package main

import "fmt"
import "strings"

func convert(s string, numRows int) string {
	t := strings.Split(s, "") 
	len := len(t)

	if ( len <= 1 || numRows <= 1 ) {
		return s;
	}

	steps, result := 2 * numRows - 2, ""

	for i := 0; i < numRows; i++ {
		for j := i; j < len; j+=steps {
			result += t[j];
			if i != 0 && i != numRows - 1 && (j + steps - 2 * i) < len {
                result += t[j + steps - 2 * i];
            }
		}
	}


	fmt.Println(result)

	return result
}