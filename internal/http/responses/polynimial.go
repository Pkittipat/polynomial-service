package responses

type Calculate struct {
	IsPolynomial bool `json:"is_polynomial"`
}

func NewCalculate(isPolynomial bool) *Calculate {
	return &Calculate{
		IsPolynomial: isPolynomial,
	}
}

type DatasetPayload struct {
	Dataset []int `json:"dataset"`
}

func NewDatasetPayload(dataset []int) *DatasetPayload {
	return &DatasetPayload{
		Dataset: dataset,
	}
}
