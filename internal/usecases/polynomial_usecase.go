package usecases

import (
	"strconv"

	"github.com/Pkittipat/polynomial-service/internal/http/responses"
)

type PolynomialUsecase interface {
	Calculate(x, y, z int) error
}

type polynomialUsecase struct {
}

var (
	dataSet = []string{"1", "x", "8", "17", "y", "z", "78", "113"}
)

func NewPolynomialUsecase() PolynomialUsecase {
	return &polynomialUsecase{}
}

func (p *polynomialUsecase) Calculate(x, y, z int) error {
	dataset := prePareDataSet(x, y, z)
	dataset = calculate(dataset, 0)
	isPolynimial := isPolynomial(dataset)
	if !isPolynimial {
		return responses.ErrInvalidPolynomial
	}
	return nil
}

func prePareDataSet(x, y, z int) (result []int) {
	for _, val := range dataSet {
		switch val {
		case "x":
			result = append(result, x)
		case "y":
			result = append(result, y)
		case "z":
			result = append(result, z)
		default:
			valInt, _ := strconv.Atoi(val)
			result = append(result, valInt)
		}
	}
	return result
}

func calculate(sequence []int, coefficient int) (_sequence []int) {
	for i, val := range sequence {
		if i+1 == len(sequence) {
			break
		}

		sub := sequence[i+1] - val
		_sequence = append(_sequence, sub)
	}

	coefficient += 1
	if !isPolynomial(_sequence) && len(_sequence) > 1 && coefficient < 3 {
		return calculate(_sequence, coefficient)
	}
	return
}

func isPolynomial(sequence []int) bool {
	// Check that the sequence has at least two elements
	if len(sequence) < 2 {
		return false
	}

	// Check that the difference between consecutive elements is constant
	for i, val := range sequence {
		if i+1 == len(sequence) {
			break
		}

		if val != sequence[i+1] {
			return false
		}
	}

	return true
}
