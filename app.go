package main

import (
	"fmt"

	"go.ir/filemanager"
	"go.ir/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		fm := filemanager.New("./storage/prices.txt", fmt.Sprintf("./storage/result_%v.json", taxRate))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index, _ := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
		}
	}
}
