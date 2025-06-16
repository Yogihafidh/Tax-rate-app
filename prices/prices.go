package prices

import (
	"fmt"
	"math"

	"taxrate.com/tax/conversion"
	"taxrate.com/tax/iomanager"
)

type TaxIncludedPricesJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_included_prices"`
}

// Constructor functon to create a new TaxIncludedPricesJob instance with input price and tax rate
func NewTaxIncludedPricesJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   iom,
	}
}

// Method Receiver for TaxIncludedPricesJob struct
func (job *TaxIncludedPricesJob) Process() error {
	// Load data from file into InputPrices
	err := job.LoadData()
	if err != nil {
		return err
	}

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
	return job.IOManager.WriteResult(job)
}

func (job *TaxIncludedPricesJob) LoadData() error {
	// Read lines from file and store in lines slice
	lines, err := job.IOManager.ReadFile()
	if err != nil {
		return err
	}

	// Convert lines slice with type string to slice with float64 and store in InputPrices
	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		return err
	}

	// Assign prices slice to InputPrices field of the job
	job.InputPrices = prices
	return nil
}
