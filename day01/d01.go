package day01

// PartOne counts the number of times a depth measurement increases from the previous measurement.
// (There is no measurement before the first measurement.)
func PartOne(depths []uint) uint {
	var count uint = 0
	for i, depth := range depths[1:] {
		if depth > depths[i] {
			count += 1
		}
	}

	return count
}
