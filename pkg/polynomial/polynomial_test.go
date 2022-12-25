package polynomial

import (
	"testing"
)

func TestPolynomial(t *testing.T) {
	testcases := []struct {
		name string
		in   []string
		want []int
	}{
		{
			name: "",
			in:   []string{"1", "x", "8", "17", "y", "z", "78", "113"},
			want: []int{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewPolynomial(tc.in)
			// result := traverseDoublyLinkedList(p.Root)
			findMissingValue(p.Root)
		})
	}
}

func toIntPtr(v int) *int {
	return &v
}

func TestFindNextValue(t *testing.T) {
	dataset := []string{"1", "x", "8", "17", "y", "z", "78", "113"}
	polynomial := NewPolynomial(dataset)
	slice := traverseDoublyLinkedList(polynomial.Root)
	testcases := []struct {
		name  string
		index int
		want  *int
	}{
		{name: "", index: 1, want: toIntPtr(8)},
		{name: "", index: 4, want: toIntPtr(78)},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := findNextValue(slice[tc.index])
			if tc.want == nil && got != nil {
				t.Errorf("expected: %v got: %v", tc.want, got)
			}
			if tc.want != nil && got != nil {
				if *tc.want != *got {
					t.Errorf("expected: %v got: %v", tc.want, got)
				}
			}
		})
	}
}

func TestFindPrevValue(t *testing.T) {
	dataset := []string{"1", "x", "8", "17", "y", "z", "78", "113"}
	polynomial := NewPolynomial(dataset)
	slice := traverseDoublyLinkedList(polynomial.Root)
	testcases := []struct {
		name  string
		index int
		want  *int
	}{
		{name: "", index: 1, want: toIntPtr(1)},
		{name: "", index: 4, want: toIntPtr(17)},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := findPrevValue(slice[tc.index])
			if tc.want == nil && got != nil {
				t.Errorf("expected: %v got: %v", tc.want, got)
			}
			if tc.want != nil && got != nil {
				if *tc.want != *got {
					t.Errorf("expected: %v got: %v", tc.want, got)
				}
			}
		})
	}
}

func TestFindDifferenceNumber(t *testing.T) {
	testcases := []struct {
		name string
		arg  []int
		want []int
	}{
		{
			name: "",
			arg:  []int{1, 3, 8, 17, 31, 51, 78, 113},
			want: []int{1, 1, 1, 1, 1},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := findDifferenceNumber(tc.arg, 0)
			if len(tc.want) != len(got) {
				t.Errorf("expected %v got %v", len(tc.want), len(got))
			}
		})
	}
}
