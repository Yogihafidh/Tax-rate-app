package cmdmanager

import "fmt"

type CMDManager struct{}

func New() CMDManager {
	return CMDManager{}
}

func (cmd CMDManager) ReadFile() ([]string, error) {
	fmt.Println("Please enter your price. Confirm every price with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scanln(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}
	return prices, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}
