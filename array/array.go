package array

func Sum(arr [5]int) int {

	var s int = 0
	for _, a := range arr {
		s += a
	}

	return s
}
