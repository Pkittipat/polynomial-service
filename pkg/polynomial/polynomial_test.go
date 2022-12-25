package polynomial

import "testing"

func TestFindNextValue(t *testing.T) {
	polynomial := NewPolynomial()
	slice := traverseDoublyLinkedList(polynomial.GetRoot())
	testcases := []struct {
		name  string
		index int
		want  int
	}{
		{name: "case_index_one", index: 1, want: 8},
		{name: "case_index_four", index: 4, want: 78},
		{name: "case_index_five", index: 5, want: 78},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := findNextValue(slice[tc.index])
			if tc.want != got {
				t.Errorf("expected: %v got: %v", tc.want, got)
			}
		})
	}
}

func TestFindPrevValue(t *testing.T) {
	polynomial := NewPolynomial()
	slice := traverseDoublyLinkedList(polynomial.GetRoot())
	testcases := []struct {
		name  string
		index int
		want  int
	}{
		{name: "index_one", index: 1, want: 1},
		{name: "index_four", index: 4, want: 17},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := findPrevValue(slice[tc.index])
			if tc.want != got {
				t.Errorf("expected: %v got: %v", tc.want, got)
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
			name: "case_polynomial_is_one",
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

func TestIsPolynomial(t *testing.T) {
	testcases := []struct {
		name string
		arg  []int
		want bool
	}{
		{
			name: "case_not_related",
			arg:  []int{1, 2, 1, 1},
			want: false,
		},
		{
			name: "case_related",
			arg:  []int{1, 1, 1, 1},
			want: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := isPolynomial(tc.arg)
			if tc.want != got {
				t.Errorf("expexted: %v, accepted: %v", tc.want, got)
			}
		})
	}
}
