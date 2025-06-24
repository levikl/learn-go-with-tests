package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slicesOfNumbers ...[]int) []int {
	var sums []int
	for _, numbers := range slicesOfNumbers {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
