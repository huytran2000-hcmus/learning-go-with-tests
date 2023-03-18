package arraysandslices

func Sum(numbers []int) int {
	var sum int
	for _, i := range numbers {
		sum += i
	}

	return sum
}

func SumAll(arraysOfNumbs ...[]int) []int {
	var sums []int

	for _, numbers := range arraysOfNumbs {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(arraysOfNumbs ...[]int) []int {
	var sums []int
	for _, numbers := range arraysOfNumbs {
		if len(numbers) > 0 {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}
