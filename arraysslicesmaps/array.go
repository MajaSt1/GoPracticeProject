package arraysslicesmaps

import "fmt"

type product struct {
	title string
	id    int
	price float64
}

func ShowListExample() {
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

func BuildDynamiLists() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	updatedPrices := append(prices, 5.99, 6.99, 7.00) // create new array
	fmt.Println(updatedPrices, prices)

	prices = prices[1:] // reassign and remove feature / there is no built in remove function

	discountPrices := []float64{10.99, 80.99, 70.00}
	prices = append(prices, discountPrices...) // pull the elements from list
}

func Exercise() {
	hobbies := [3]string{"drums", "piano", "violin"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	mainHobbies := hobbies[0:2]
	fmt.Println(mainHobbies)

	fmt.Println(cap(mainHobbies), mainHobbies[1:3]) // expanding explicitly from parent array

	goals := []string{"goal1", "goal2"}
	goals[1] = "goal3"
	goals = append(goals, "goal4") // create new array and assign to existing one
	fmt.Println(goals)

	products := []product{{"book", 0, 1.99}, {"door", 1, 109.99}}
	products = append(products, product{"pen", 2, 3.99})
	fmt.Println(products)
}

// make memory more efficient
func SpecialListMakeFunc() {
	userNames := make([]string, 2, 5) // 2- initial length of slice, 5- capacity = how much space will be allocated BTS by go, when append >5 elements it will have to allocate new space

	//userNames := []string{} - when userNames[0] = "Julie" it will result with error
	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)
}
