package projectpricecalculator

import (
	"fmt"

	"example.com/note/project-price_calculator/conversion"
	"example.com/note/project-price_calculator/filemanager"
)

func ShowPriceCalculator() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		priceJob := NewTaxIncludedPriceJob(taxRate)
		priceJob.process()
	}
}

func (job *TaxIncludedPriceJob) loadData() {
	lines, err := filemanager.Readlines("project-price_calculator/prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	// overwritten input prices
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) process() {
	job.loadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob { // create this value only once in memory and we share address to that value
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
