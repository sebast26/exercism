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

func (r Record) valid() bool {
	if r.ID == 0 && r.Parent != 0 {
		return false
	}
	return true
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	var nMap = make(map[int]*Node, len(records))
	for _, record := range records {
		if !record.valid() {
			return nil, ErrInvalidInput
		}
		nMap[record.ID] = &Node{
			ID:       record.ID,
			Children: make([]*Node, 0),
		}
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	var root *Node
	for _, record := range records {
		parent, ok := nMap[record.Parent]
		if !ok {
			continue
		}
		if record.ID == 0 {
			root = nMap[record.ID]
			continue
		}
		parent.Children = append(parent.Children, nMap[record.ID])
	}

	//if root == nil {
	//	return nil, errors.New("")
	//}
	return root, nil
}
