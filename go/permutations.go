package kata

func Permutations(s string) []string {
	sArr := []string{}
	sMap := make(map[string]bool)
	if len(s) == 1 {
		sArr = append(sArr, s)
		return sArr
	} else {
		for i, v := range s {
			remain := s[:i] + s[i+1:]
			for _, permutate := range Permutations(remain) {
				permutation := string(v) + permutate
				if !sMap[permutation] {
					sArr = append(sArr, permutation)
					sMap[permutation] = true
				}
			}
		}
	}
	return sArr
}
