package kata

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	} else {
		time := float64(g*3600) / float64(v2-v1)
		h := int(time) / 3600
		mn := (int(time) % 3600) / 60
		s := int(time) % 60
		timeRequired := [3]int{h, mn, s}
		return timeRequired
	}
}
