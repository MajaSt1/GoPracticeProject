package functions

import "fmt"

func ShowRecursionExample() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1) // factorial(number-1) return always previous execution results
	// result := 1
	// for i := 1; i < number; i++ {
	// 	result = number * i
	// }
	// return result
}

// factorial of 5 => 5 * 4 * 3 * 2 * 1 => 120
// 5 * 4 = 20
// 20 * 3 = 60
// 60 * 2 = 120
// 120 * 1 = 120
