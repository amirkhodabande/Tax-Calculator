package prices

import (
	"fmt"

	"go.ir/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]string
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadFloatFromFile("./storage/prices.txt")

	if err != nil {
		fmt.Println(err)
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
	job.TaxIncludedPrice = result

	filemanager.WriteJson(fmt.Sprintf("./storage/result_%v.json", job.TaxRate), job)

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
