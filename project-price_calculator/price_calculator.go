package projectpricecalculator

import (
	"fmt"

	"example.com/note/project-price_calculator/conversion"
	"example.com/note/project-price_calculator/filemanager"
)

func ShowPriceCalculator() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		fm := filemanager.New("project-price_calculator/prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.process()
	}
}

func (job *TaxIncludedPriceJob) loadData() {

	lines, err := job.IOManager.Readlines()
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
	job.IOManager.WriteResult(job)
}

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob { // create this value only once in memory and we share address to that value
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
