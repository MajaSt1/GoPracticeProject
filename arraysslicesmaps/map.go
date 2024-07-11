package arraysslicesmaps

import "fmt"

func ShowMapExample() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])
	websites["Linkedin"] = "https://linkedin.com"

	delete(websites, "Google")
	fmt.Println(websites)
}

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func SpecialMapMakeFunc() {
	courseRatings := make(floatMap, 3) // 3 - pre-allocate memory

	// go will have realocate new memory if we will add new items to courseRatings := map[string]float64{}
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8

	courseRatings.output()

	for key, value := range courseRatings {
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)
	}
}
