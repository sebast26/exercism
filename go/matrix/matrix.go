package matrix

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	m [][]int
}

func New(s string) (*Matrix, error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	var m [][]int
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(strings.TrimSpace(line), " ")
		if len(items) == 0 {
			return nil, fmt.Errorf("uneven rows")
		}
		numbers := make([]int, 0, len(items))
		for _, item := range items {
			num, err := strconv.Atoi(item)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, num)
		}
		m = append(m, numbers)
	}

	rowItems := len(m[0])
	for _, row := range m {
		if rowItems != len(row) {
			return nil, fmt.Errorf("uneven rows")
		}
	}

	return &Matrix{m: m}, scanner.Err()
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	panic("Please implement the Cols function")
}

func (m *Matrix) Rows() [][]int {
	panic("Please implement the Rows function")
}

func (m *Matrix) Set(row, col, val int) bool {
	panic("Please implement the Set function")
}
