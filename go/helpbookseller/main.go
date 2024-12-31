package main

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {

	var results []string
	for _, v := range listCat {
		sum := 0
		for _, c := range listArt {
			if strings.HasPrefix(c, v) {
				for i, char := range c {
					if char == ' ' {
						n, _ := strconv.Atoi(c[i+1:])
						sum += n
					}
				}
			}
		}
		results = append(results, fmt.Sprintf("(%s : %d)", v, sum))
	}
	return strings.Join(results, " - ")
}

func main() {
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}

	fmt.Println(StockList(b, c))
}
