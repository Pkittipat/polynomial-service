package polynomial

import (
	"strconv"
)

const (
	X = "X"
	Y = "Y"
	Z = "Z"
)

var (
	defaultDataset = []string{
		"1",
		X,
		"8",
		"17",
		Y,
		Z,
		"78",
		"113",
	}
)

type polynomial struct {
	Root *Node
}

func NewPolynomial(
	dataset []string,
) *polynomial {
	if len(dataset) <= 0 {
		return nil
	}
	root := createDoublyLinkedList(dataset)
	return &polynomial{
		Root: root,
	}
}

func createDoublyLinkedList(slice []string) *Node {
	var head, tail *Node
	for i, v := range slice {
		n := &Node{}
		_v, err := strconv.Atoi(v)
		if err == nil {
			n.Value = &_v
		}
		if i == 0 {
			head = n
			tail = n
		} else {
			tail.Next = n
			n.Prev = tail
			tail = n
		}
	}
	return head
}

func traverseDoublyLinkedList(head *Node) []*Node {
	var result []*Node
	current := head
	for current != nil {
		result = append(result, current)
		current = current.Next
	}
	return result
}

func findMissingValue(head *Node) {
	slice := traverseDoublyLinkedList(head)
	missingIndexs := []int{}
	for index, node := range slice {
		if node.Value == nil {
			missingIndexs = append(missingIndexs, index)
		}
	}

	var expectArgs [][]int
	for _, missingIndex := range missingIndexs {
		prevVal := findPrevValue(slice[missingIndex])
		nextVal := findNextValue(slice[missingIndex])
		var expectArr []int
		for num := *prevVal; num <= *nextVal; num++ {
			expectArr = append(expectArr, num)
		}
		expectArgs = append(expectArgs, expectArr)
	}

	// specify x represent first value.
	for _, x := range expectArgs[0] {
		// specify y represent second value
		for _, y := range expectArgs[1] {
			for _, z := range expectArgs[2] {
				slice[missingIndexs[0]].Value = &x
				slice[missingIndexs[1]].Value = &y
				slice[missingIndexs[2]].Value = &z
			}
		}
	}
	return
}

func findDifferenceNumber(dataset []int, degree int) (sequence []int) {
	for index, value := range dataset {
		if index+1 == len(dataset) {
			break
		}
		diff := dataset[index+1] - value
		sequence = append(sequence, diff)
	}

	degree += 1
	if degree >= 3 {
		return
	}

	if len(sequence) < 1 {
		return
	}

	if isPolynomial(sequence) {
		return
	}

	return findDifferenceNumber(sequence, degree)
}

func isPolynomial(sequence []int) bool {
	if len(sequence) < 2 {
		return false
	}

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

func findPrevValue(head *Node) (prevValue *int) {
	if head == nil {
		return
	}
	if head.Prev == nil {
		return
	}
	if head.Prev.Value == nil {
		prevValue = findPrevValue(head.Prev)
	} else {
		prevValue = head.Prev.Value
	}
	return prevValue
}

func findNextValue(head *Node) (nextValue *int) {
	if head == nil {
		return
	}

	if head.Next == nil {
		return
	}

	nextValue = head.Next.Value
	if nextValue == nil {
		nextValue = findNextValue(head.Next)
	}
	return
}

// func (p *polynomial) Initialize(preDataset []string) {
// 	p.preDataset = defaultDataset
// 	if len(preDataset) > 0 {
// 		p.preDataset = preDataset
// 	}

// 	return
// }

// func (p *polynomial) findMissingValue() {
// 	indexMissingValue := []int{}
// 	for idx, val := range p.preDataset {
// 		_, err := strconv.Atoi(val)
// 		if err != nil {
// 			indexMissingValue = append(indexMissingValue, idx)
// 		}
// 	}

// 	return
// }

func (p *polynomial) pareDataset() {
	// for _, value := range p.PreDataset {
	// 	switch value {
	// 	case X:
	// 		p.dataset = append(p.dataset, value)
	// 	case Y:
	// 		p.dataset = append(p.dataset, value)
	// 	case Z:
	// 		p.dataset = append(p.dataset, value)
	// 	default:
	// 		tmp, _ := strconv.Atoi(value)
	// 		p.dataset = append(p.dataset, tmp)
	// 	}
	// }
}

// func (p *polynomial) caludate(coefficient int) int {
// 	result := math.Pow(float64(coefficient), 3)*p.args["a"] +
// 		math.Pow(float64(coefficient), 2) +
// 		float64(coefficient)*p.args["c"] +
// 		p.args["d"]
// 	return int(result)
// }
