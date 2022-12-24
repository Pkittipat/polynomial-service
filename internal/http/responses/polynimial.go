package responses

type Calculate struct {
	IsPolynomial bool `json:"is_polynomial"`
}

func NewCalculate(isPolynomial bool) *Calculate {
	return &Calculate{
		IsPolynomial: isPolynomial,
	}
}
