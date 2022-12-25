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

type Polynomial interface {
	RandomArg() (datasetRes []int)
	GetRoot() *Node
	ProveXYZ(x, y, z int) ([]int, bool)
}

type polynomial struct {
	Root *Node
}

func NewPolynomial() Polynomial {
	if len(defaultDataset) <= 0 {
		return nil
	}
	root := createDoublyLinkedList()
	return &polynomial{
		Root: root,
	}
}

func (p *polynomial) GetRoot() *Node {
	return p.Root
}

func (p *polynomial) ProveXYZ(x, y, z int) ([]int, bool) {
	dataset := prepareDataset(p.Root, x, y, z)
	sliceOfDiff := findDifferenceNumber(dataset, 0)
	return dataset, isPolynomial(sliceOfDiff)
}

func (p *polynomial) RandomArg() (datasetRes []int) {
	indexRange := createXYZIndexRange(p.Root)

	for _, x := range indexRange[0] {
		for _, y := range indexRange[1] {
			for _, z := range indexRange[2] {
				dataset, isPolynomial := p.ProveXYZ(x, y, z)
				if isPolynomial {
					return dataset
				}
			}
		}
	}

	return
}

func createXYZIndexRange(head *Node) (indexRange [][]int) {
	slice := traverseDoublyLinkedList(head)
	indexMissingValues := findIndexMissingValues(head)

	for _, missingIndex := range indexMissingValues {
		prevVal := findPrevValue(slice[missingIndex])
		nextVal := findNextValue(slice[missingIndex])
		var scope []int
		for num := prevVal; num <= nextVal; num++ {
			scope = append(scope, num)
		}
		indexRange = append(indexRange, scope)
	}
	return indexRange
}

func prepareDataset(head *Node, x, y, z int) (result []int) {
	current := head
	for current != nil {
		switch current.Value {
		case X:
			result = append(result, x)
		case Y:
			result = append(result, y)
		case Z:
			result = append(result, z)
		default:
			val, _ := strconv.Atoi(current.Value)
			result = append(result, val)
		}
		current = current.Next
	}
	return result
}

func createDoublyLinkedList() *Node {
	var head, tail *Node
	for i, v := range defaultDataset {
		n := &Node{Value: v}
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

func findIndexMissingValues(head *Node) (slice []int) {
	current := head
	index := 0
	for current != nil {
		switch current.Value {
		case X, Y, Z:
			slice = append(slice, index)
		}
		index += 1
		current = current.Next
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

func findPrevValue(head *Node) (prevValue int) {
	if head == nil {
		return
	}
	if head.Prev == nil {
		return
	}
	switch head.Prev.Value {
	case X, Y, Z:
		prevValue = findPrevValue(head.Prev)
	default:
		val, _ := strconv.Atoi(head.Prev.Value)
		prevValue = val
	}
	return prevValue
}

func findNextValue(head *Node) (nextValue int) {
	if head == nil {
		return
	}

	if head.Next == nil {
		return
	}

	switch head.Next.Value {
	case X, Y, Z:
		nextValue = findNextValue(head.Next)
	default:
		val, _ := strconv.Atoi(head.Next.Value)
		nextValue = val
	}
	return
}
