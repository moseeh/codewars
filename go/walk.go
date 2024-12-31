package kata

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}
	var countn, counts, countw, counte int
	for _, d := range walk {
		if d == 'n' {
			countn++
		}
		if d == 's' {
			counts++
		}
		if d == 'w' {
			countw++
		}
		if d == 'e' {
			counte++
		}
	}
	if countn != counts {
		return false
	}
	if counte != countw {
		return false
	}
	return true
}
