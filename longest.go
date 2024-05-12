package kata

// returns the longest k consecutive words from an array
func LongestConsec(strarr []string, k int) string {
	max := 0
	s := ""
	startindex := 0

	if len(strarr) == 0 || k > len(strarr) {
		return ""
	}
	for i := 0; i <= len(strarr)-k; i++ {
		lenword := 0
		for j := i; j < i+k; j++ {
			lenword += len(strarr[j])
		}
		if lenword > max {
			max = lenword
			startindex = i
		}
	}
	for i := startindex; i < startindex+k; i++ {
		s += strarr[i]
	}
	return s
}
