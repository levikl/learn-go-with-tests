package main

func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		}

		tail := x[1:]
		return append(acc, Sum(tail))
	}

	return Reduce(numbers, sumTail, []int{})
}
