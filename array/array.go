package array

func Sum(arr []int) int {

	var s int = 0
	for _, a := range arr {
		s += a
	}

	return s
}

func SumAll(numbersToSum ...[]int) []int {

	length := len(numbersToSum)
	result := make([]int, length)

	for i, numbers := range numbersToSum {
		s := 0
		for _, num := range numbers {
			s += num
		}
		result[i] = s
	}

	return result
}
