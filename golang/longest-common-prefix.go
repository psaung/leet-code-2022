package main


func longestCommonPrefix(strs []string) string {
	lens, result := len(strs), ""

	if lens <= 1 {
		return strs[0]
	}

	first := strs[0]

	for i, firstLen := 0, len(first); i < firstLen; i++ {
		for j := 1; j < lens; j++ {
			if i > len(strs[j]) - 1 || strs[0][i] != strs[j][i] {
				return result;
			}
		}

		result = result + string(strs[0][i])
	}

	return result;
}