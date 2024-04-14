package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("./storage/prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []float64

	for scanner.Scan() {
		floatPrice, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			fmt.Println(err)
			file.Close()
			return
		}

		lines = append(lines, floatPrice)
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	job.InputPrices = lines
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)

		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}