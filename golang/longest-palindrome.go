package main


func substring(s string, start int, end int) string {
    start_str_idx := 0
    i := 0
    for j := range s {
        if i == start {
            start_str_idx = j
        }
        if i == end {
            return s[start_str_idx:j]
        }
        i++
    }
    return s[start_str_idx:]
}

func longestPalindrome(s string) string {
	ll, rr, len := 0, 0, len(s) 

	for i:= 0; i < len; i++ {
		for j := i; j <= i+1; j++ {
 			l, r := i, j
			for l >= 0 && r < len && s[l] == s[r] {
				if(r-l+1 > rr-ll+1) {
					ll = l
					rr = r
				}
				l--;
				r++;
			}
		}
	} 

	return substring(s, ll, rr + 1);
}