package functions

import "fmt"

// anonymous func - feature that allows you to define a function just in time when you need it instead of in advance
func ShowAnonymousFunc() {
	numbers := []int{1, 2, 3}

	transformed := anonymousTransformNumbers(&numbers, func(number int) int {
		return number * 2
	}) // its not the type of a function, we don't need to reuse and name this logic

	fmt.Println(transformed)
}

func anonymousTransformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
