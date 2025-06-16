package prices

import (
	"fmt"
	"math"

	"taxrate.com/tax/conversion"
	"taxrate.com/tax/filemanager"
)

type TaxIncludedPricesJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// Constructor functon to create a new TaxIncludedPricesJob instance with input price and tax rate
func NewTaxIncludedPricesJob(taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

// Method Receiver for TaxIncludedPricesJob struct
func (job *TaxIncludedPricesJob) Process() {
	// Load data from file into InputPrices
	job.LoadData()

	// Create map without initial values using make function
	result := make(map[string]float64)

	// Loop for looping prices. 10, 20, 30
	for _, price := range job.InputPrices {
		// Calculate price after adding tax and resul store in the priceIndex position in taxIncludedPrices
		taxIncludePrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.f", price)] = math.Round(taxIncludePrice*100) / 100
	}

	// Assign result map to TaxIncludedPrices field of the job
	job.TaxIncludedPrices = result
	fmt.Println(result)
}

func (job *TaxIncludedPricesJob) LoadData() {
	// Read lines from file and store in lines slice
	lines, err := filemanager.ReadFile("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert lines slice with type string to slice with float64 and store in InputPrices
	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Assign prices slice to InputPrices field of the job
	job.InputPrices = prices
}
