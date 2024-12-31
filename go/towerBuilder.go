package kata

import "strings"

func TowerBuilder(nFloors int) []string {
	tower := make([]string, nFloors)
	for i := 0; i < nFloors; i++ {
		spaces := strings.Repeat(" ", nFloors-i-1)
		stars := strings.Repeat("*", 2*i+1)
		tower[i] = spaces + stars + spaces
	}
	return tower
}
