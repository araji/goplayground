package tree


import (
	
	"errors"
	"sort"
)

type Record struct {
	ID,Parent    int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records)  == 0 {
		return nil,nil
	}
	
	sort.SliceStable(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	
	root := &Node{ID: records[0].ID}
	
	if records[0].ID != records[0].Parent || records[0].ID != 0 {
		return root, errors.New("invalid root node")
	}

	maxID := len(records) - 1	
	nodes := map[int]*Node{root.ID: root}

	for _, record := range records[1:] {
		if record.ID <= record.Parent || record.ID > maxID {
			return root, errors.New("invalid ID")
		}
		
		parent, ok := nodes[record.Parent]
		if !ok {
			return root, errors.New("invalid Parent")
		}
		
		node := &Node{ID: record.ID}

		if _,ok := nodes[record.ID]; ok {
			return root, errors.New("Duplicate Key")
		} else {
			nodes[record.ID] = node
			parent.Children = append(parent.Children, node)
		}

	}
	return root,nil
}

