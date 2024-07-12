package functions

import "fmt"

func ShowExampleOfClosure() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2) // factory function pattern
	triple := createTransformer(3)

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled, tripled)
}

// closure - value will locked in and be available at any point in the future
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
