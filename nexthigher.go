package kata

func NextHigher(n int) int {
	smallest := n & -n     // rightmost 1-bit
	ripple := n + smallest // add the smallest 1-bit to x

	ones := ((ripple ^ n) / smallest) >> 2 // right bits rearranged

	// Step 2: Construct the next higher number
	nextHigher := ripple | ones // combine the results

	return nextHigher
}
