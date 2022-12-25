package responses

type GuessPayload struct {
	IsPolynomial bool `json:"is_polynomial"`
}

func NewGuessPayload(isPolynomial bool) *GuessPayload {
	return &GuessPayload{
		IsPolynomial: isPolynomial,
	}
}

type DatasetPayload struct {
	Dataset []int `json:"dataset"`
	Degree  int   `json:"degree"`
	Range   int   `json:"range"`
}

func NewDatasetPayload(dataset []int, degree, diff int) *DatasetPayload {
	return &DatasetPayload{
		Dataset: dataset,
		Degree:  degree,
		Range:   diff,
	}
}
