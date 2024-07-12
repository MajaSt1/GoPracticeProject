package functions

import "fmt"

func ShowVariadicFunc() {
	// numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, 40, -5)

	fmt.Println(sum)
}

// variadic - function that can take any amount of parameters (dynamic) - BTS it will merge within a slice for you
// []int - slice
func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val // sum = sum + val
	}

	return sum
}
