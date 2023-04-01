package generics

func Sum(numbers []int) int {
	add := func(accumulated, val int) int {
		return accumulated + val
	}

	sum := Reduce(numbers, add, 0)

	return sum
}

func SumAll(arraysOfNumbs ...[]int) []int {
	addAll := func(accumulated, val []int) []int {
		return append(accumulated, Sum(val))
	}
	sums := Reduce(arraysOfNumbs, addAll, []int{})

	return sums
}

func SumAllTails(arraysOfNumbs ...[]int) []int {
	addAllTails := func(accumulated, val []int) []int {
		if len(val) > 0 {
			tail := val[1:]
			return append(accumulated, Sum(tail))
		} else {
			return append(accumulated, 0)
		}
	}

	sums := Reduce(arraysOfNumbs, addAllTails, []int{})

	return sums
}

func Reduce[T, R any](collection []T, accumulator func(R, T) R, initalVal R) R {
	result := initalVal
	for _, a := range collection {
		result = accumulator(result, a)
	}

	return result
}
