package main

import (
	"taxrate.com/tax/prices"
)

func main() {
	taxRate := []float64{0, 0.07, 0.1, 0.15}

	// Loop for looping tax rates
	for _, taxRate := range taxRate {
		// Create struct instance of TaxIncludedPricesJob with tax rate as parameter
		pricesJob := prices.NewTaxIncludedPricesJob(taxRate)
		// Call Process method to calculate tax included prices
		pricesJob.Process()
	}

}
