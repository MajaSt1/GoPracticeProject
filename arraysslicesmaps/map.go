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

func SpecialMapMakeFunc() {
	courseRatings := make(map[string]float64, 3) // 3 - pre-allocate memory

	// go will have realocate new memory if we will add new items to courseRatings := map[string]float64{}
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8

	fmt.Println(courseRatings)
}
