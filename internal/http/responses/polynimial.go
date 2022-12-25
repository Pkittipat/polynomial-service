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
}

func NewDatasetPayload(dataset []int) *DatasetPayload {
	return &DatasetPayload{
		Dataset: dataset,
	}
}
