package projectpricecalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ShowPriceCalculator() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		priceJob := NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) loadData() {
	file, err := os.Open("project-price_calculator/prices.txt")

	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading file content failed!")
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Converting price to float failed!")
			fmt.Println(err)
			file.Close()
			return
		}

		prices[lineIndex] = floatPrice
	}

	// overwritten input prices
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.loadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob { // create this value only once in memory and we share address to that value
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
