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

func main() {
	Sum([5]int{1, 2, 3, 4, 5})
}
