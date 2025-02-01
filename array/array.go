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
		result[i] = Sum(numbers)
	}

	return result
}
