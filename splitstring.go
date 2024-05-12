package kata

func Solution(str string) []string {
	s := ""
	sArr := []string{}
	for i, v := range str {
		if i%2 == 0 {
			s += string(v)
		} else {
			s += string(v)
			sArr = append(sArr, s)
			s = ""
		}
		if i == len(str)-1 && len(str)%2 == 1 {
			s += "_"
			sArr = append(sArr, s)
		}
	}
	return sArr
}
