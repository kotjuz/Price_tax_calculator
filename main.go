package main

import (
	"fmt"

	"example.com/calculator/filemanager"
	"example.com/calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		// cm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// priceJob2 := prices.NewTaxIncludedPriceJob(cm, taxRate)
		err := priceJob.Process()
		// priceJob2.Process()
		if err != nil {
			fmt.Println("Could not process job ", err)
		}
	}

}
