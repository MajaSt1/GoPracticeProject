package arraysslicesmaps

import "fmt"

func ShowExample() {
	var productNames [4]string = [4]string{"A book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	productNames[2] = "A carpet"
	fmt.Println(prices[0])
	fmt.Println(productNames)

	//slices
	featuredPrices := prices[1:3] // [:3] [1:] window to the parent array
	featuredPrices[0] = 199.99
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(prices)                                         // [0] 9.99 = 199.99 will be overwritten in parent
	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // cap - counts the original array

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // there is always more content available because its a window /slice in memory of parent array
}
