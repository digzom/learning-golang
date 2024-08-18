package sum

func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func Sum2(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// this creates a slice object
	// sums := make([]int, lengthOfNumbers)
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum2(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) >= 1 {
			tail := numbers[1:]
			sums = append(sums, Sum2(tail))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}

func main() {
	Sum([5]int{1, 2, 3, 4, 5})
}
