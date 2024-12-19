package arrayandslices

func Sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total += number
	}

	return total
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

//difference between length and capacity in a slice
