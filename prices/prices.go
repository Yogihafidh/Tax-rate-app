package prices

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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
		result[fmt.Sprintf("%.f", price)] = math.Round(taxIncludePrice * 100) /100
	}

	// Assign result map to TaxIncludedPrices field of the job
	job.TaxIncludedPrices = result
	fmt.Println(result)
}

func (job *TaxIncludedPricesJob) LoadData() {
	// If successful, methods on the returned file can be used for reading. Method return io.Reader interface
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could not open file!", err)
		return
	}

	// New scanner reciving input value with type io.Reader interface
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Error handling for scanner
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		file.Close()
		return
	}

	// Convert lines slice with type string to slice with float64 and store in InputPrices
	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64) // Convert string to float64

		if err != nil { // Error handling for conversion
			fmt.Printf("Error converting line %d to float: %v\n", lineIndex, err)
			file.Close()
			return
		}

		prices[lineIndex] = floatPrice // Store converted float64 value in prices slice

	}

	// Assign prices slice to InputPrices field of the job
	job.InputPrices = prices
}
