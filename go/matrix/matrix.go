package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	m [][]int
}

var ErrInvalidInput = errors.New("invalid input")

func New(s string) (*Matrix, error) {
	lines := strings.Split(s, "\n")

	var m [][]int
	for _, line := range lines {
		items := strings.Split(strings.TrimSpace(line), " ")
		if len(items) == 0 {
			return nil, ErrInvalidInput
		}
		numbers := make([]int, 0, len(items))
		for _, item := range items {
			num, err := strconv.Atoi(item)
			if err != nil {
				return nil, ErrInvalidInput
			}
			numbers = append(numbers, num)
		}
		m = append(m, numbers)
	}

	rowItems := len(m[0])
	for _, row := range m {
		if rowItems != len(row) {
			return nil, ErrInvalidInput
		}
	}

	return &Matrix{m: m}, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	ret := make([][]int, 0, len(m.m[0]))
	for i := 0; i < len(m.m[0]); i++ {
		col := make([]int, 0, len(m.m))
		for j := 0; j < len(m.m); j++ {
			col = append(col, m.m[j][i])
		}
		ret = append(ret, col)
	}
	return ret
}

func (m *Matrix) Rows() [][]int {
	ret := make([][]int, 0, len(m.m))
	for _, r := range m.m {
		cp := make([]int, len(r))
		copy(cp, r)
		ret = append(ret, cp)
	}
	return ret
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= len(m.m[0]) || col >= len(m.m) {
		return false
	}
	m.m[row][col] = val
	return true
}
