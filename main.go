package main

import (
	"fmt"

	// "taxrate.com/tax/cmdmanager"
	"taxrate.com/tax/filemanager"
	"taxrate.com/tax/prices"
)

func main() {
	taxRate := []float64{0, 0.07, 0.1, 0.15}

	// Loop for looping tax rates
	for _, taxRate := range taxRate {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		// Create struct instance of TaxIncludedPricesJob with tax rate as parameter
		pricesJob := prices.NewTaxIncludedPricesJob(fm, taxRate)
		// Call Process method to calculate tax included prices
		err := pricesJob.Process()
		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}

}
