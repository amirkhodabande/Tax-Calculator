package main

import (
	"fmt"

	"go.ir/filemanager"
	"go.ir/prices"
)

func main() {
	taxRates := []float64{0., 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("./storage/prices11.txt", fmt.Sprintf("./storage/result_%v.json", taxRate))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		err := priceJob.Process()

		if err != nil {
			fmt.Println(err)

			return
		}
	}
}
