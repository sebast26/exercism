package kindergarten

import (
	"bufio"
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Define the Garden type here.

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

type Garden struct {
	kindPlants map[string]string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	lines, err := readGarden(diagram, len(children))
	if err != nil {
		return nil, err
	}
	kids, err := readChildren(children)
	if err != nil {
		return nil, err
	}

	var m = map[string]string{}
	for i, child := range kids {
		m[child] = fmt.Sprintf("%s%s", lines[0][i*2:i*2+2], lines[1][i*2:i*2+2])
	}
	return &Garden{kindPlants: m}, nil
}

func readGarden(diagram string, num int) ([]string, error) {
	var lines []string
	s := bufio.NewScanner(strings.NewReader(diagram))
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	if len(lines) != 3 {
		return nil, errors.New("invalid input")
	}
	for _, l := range lines[1:3] {
		if len(l) != num*2 {
			return nil, errors.New("invalid input")
		}
		pos := strings.IndexFunc(l, func(r rune) bool {
			return r != 'G' && r != 'C' && r != 'R' && r != 'V'
		})
		if pos != -1 {
			return nil, errors.New("invalid input")
		}
	}

	return lines[1:3], nil
}

func readChildren(children []string) ([]string, error) {
	m := make(map[string]struct{}, len(children))
	for _, child := range children {
		m[child] = struct{}{}
	}

	if len(children) != len(m) {
		return nil, fmt.Errorf("duplicated child")
	}

	kids := make([]string, len(children))
	copy(kids, children)
	sort.Strings(kids)
	return kids, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := g.kindPlants[child]
	if !ok {
		return nil, false
	}

	var names []string
	for _, plant := range plants {
		switch plant {
		case 'V':
			names = append(names, "violets")
		case 'R':
			names = append(names, "radishes")
		case 'C':
			names = append(names, "clover")
		case 'G':
			names = append(names, "grass")
		}
	}
	return names, true
}
