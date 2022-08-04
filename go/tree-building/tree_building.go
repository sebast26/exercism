package tree

import (
	"errors"
	"sort"
)

var (
	ErrInvalidInput = errors.New("invalid record(s)")
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if err := checkRecords(records); err != nil {
		return nil, err
	}

	nodes := make([]*Node, len(records))
	for i, r := range records {
		nodes[i] = &Node{
			ID:       r.ID,
			Children: make([]*Node, 0),
		}
	}

	var root *Node
	for i, r := range records {
		if r.ID == 0 {
			root = nodes[r.ID]
			continue
		}
		parent := nodes[r.Parent]
		parent.Children = append(parent.Children, nodes[i])
	}

	return root, nil
}

func checkRecords(records []Record) error {
	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return ErrInvalidInput
		}
	}

	return nil
}
