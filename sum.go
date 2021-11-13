package algorithms

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	if (len(numbers) == 0) {
		return 0
	}

	sum := 0;

	for _, num := range numbers {
		sum += num
	}

	return sum
}