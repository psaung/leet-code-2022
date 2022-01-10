
package main


func intToRoman(num int) string {
	keys := [7]int{1, 5, 10, 50, 100, 500, 1000}
	values:= [7]string{"I", "V", "X", "L", "C", "D", "M"}
	
	t , result, idx, nxtIdx := num, "", 6, 6 

	for t != 0 {
		if idx > 0 && keys[idx] - keys[idx-1] == keys[idx - 1] {
			nxtIdx = idx - 2 
		} else {
			nxtIdx = idx - 1 
		}

		if t - keys[idx] >= 0 || (nxtIdx >= 0 && t < keys[idx] && t >= keys[idx] - keys[nxtIdx]) {
			
			if nxtIdx >= 0 && t < keys[idx] && t >= keys[idx] - keys[nxtIdx] {
				t = t - keys[idx] + keys[nxtIdx]
				result +=  values[nxtIdx] + values[idx] 
			} else {
				t = t - keys[idx]
				result += values[idx] 
			}
		}

		if nxtIdx >= 0 && t < keys[idx] && t >= keys[idx] - keys[nxtIdx] {
			continue
		}

		if t < keys[idx] {
			idx--;
		}
	}


	return result 
}