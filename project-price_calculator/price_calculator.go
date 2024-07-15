package projectpricecalculator

import "fmt"

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

func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob { // create this value only once in memory and we share address to that value
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
