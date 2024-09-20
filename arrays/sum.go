package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	results := make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		results[i] = Sum(numbers)
	}

	return results
}

func SumAllTail(numbersToSum ...[]int) []int {
	// sum all but not the first
	results := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			results[i] = 0
		} else {
			results[i] = Sum(numbers[1:])
		}
	}
	return results
}
