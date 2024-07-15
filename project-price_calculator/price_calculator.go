package projectpricecalculator

import (
	"errors"
	"fmt"

	// "example.com/note/project-price_calculator/cmdmanager"
	"example.com/note/project-price_calculator/conversion"
	"example.com/note/project-price_calculator/filemanager"
	"example.com/note/project-price_calculator/iomanager"
)

func ShowPriceCalculator() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("project-price_calculator/prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	// for _, errorChan := range errorChans {
	// 	<-errorChan
	// } fatal error: all goroutines are asleep - deadlock!

	for index := range taxRates {
		select { //similar to switch statement
		// the idea behind this: the case that gives us the vaule earlier wins and the other cases discarded
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}

	// for _, dondoneChan := range doneChans {
	// 	<-dondoneChan
	// }
}

func (job *TaxIncludedPriceJob) loadData() error {

	lines, err := job.IOManager.Readlines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		return err
	}

	// overwritten input prices
	job.InputPrices = prices
	return nil
}

// methods inside process are running in parallel
func (job *TaxIncludedPriceJob) process(doneChan chan bool, errorChan chan error) {
	err := job.loadData()

	//simulate error for select statement
	// errorChan <- errors.New("An error!")

	if err != nil {
		// return err
		// commented because of goroutines module, but you can use this function as goroutine and regular function
		errorChan <- err
		return // we must still cancel the execution of further functions
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
	doneChan <- true
}

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob { // create this value only once in memory and we share address to that value
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
