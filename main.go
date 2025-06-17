package main

import (
	"fmt"

	// "taxrate.com/tax/cmdmanager"
	"taxrate.com/tax/filemanager"
	"taxrate.com/tax/prices"
)

func main() {
	taxRate := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRate))
	errorChans := make([]chan error, len(taxRate))

	// Loop for looping tax rates
	for index, taxRate := range taxRate {
		// Replace slot slice
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		// Create struct instance of TaxIncludedPricesJob with tax rate as parameter
		pricesJob := prices.NewTaxIncludedPricesJob(fm, taxRate)
		// Call Process method to calculate tax included prices with goroutines
		go pricesJob.Process(doneChans[index], errorChans[index])

	}

	// Managing chanel and waiting until every chanel has emmited one value
	for index := range taxRate {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
