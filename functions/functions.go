package functions

import "fmt"

// !!!
type transformFn func(int) int

// map where the key is string and values are slices of integers
// type anotherFn func(int, []string, map[string][]int) ([]int, string)

func ShowExample() {
	numbers := []int{1, 2, 3, 4}
	// functions are just values in go
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled, tripled)

	transformFn1 := getTransformerFn(&numbers)
	transformFn2 := getTransformerFn(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformFn2)

	fmt.Println(transformedNumbers, moreTransformedNumbers)
}

// pointer to slice of integers
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFn(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}
