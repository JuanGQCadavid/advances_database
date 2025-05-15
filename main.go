package main

import (
	"log"
	"sort"
)

var (
	DB []*Student
)

type BTree struct {
	Root       *Node
	UpperLimit int
	LowerLimit int
}

type Node struct {
	Ids      []*Data
	Children []*Node
	Parent   *Node
}

type Data struct {
	Id       int
	Position int
}

type Student struct {
	Number int
	Name   string
	Gender string
	Age    float64
	City   string
}

func (tree *BTree) BalanceTree(node *Node, data *Data) {

	node.Ids = append(node.Ids, data)
	sort.Slice(node.Ids,
		func(i, j int) bool {
			return node.Ids[i].Id < node.Ids[j].Id
		},
	)

	midIndex := len(node.Ids) / 2
	midData := node.Ids[midIndex]

	right := &Node{Ids: append([]*Data{}, node.Ids[midIndex+1:]...)}
	node.Ids = append([]*Data{}, node.Ids[:midIndex]...)

	// left := &Node{Ids: append([]*Data{}, node.Ids[:midIndex]...)}
	// right := &Node{Ids: append([]*Data{}, node.Ids[midIndex+1:]...)}

	// Lets balance the pointers if we are in a mid node

	if node.Children != nil {
		// left.Children = node.Children[:midIndex+1]
		node.Children = node.Children[:midIndex+1]
		right.Children = node.Children[midIndex+1:]
	}

	// We are on the root, easy case
	if node.Parent == nil {
		log.Println("New root")
		newRoot := &Node{
			Ids: []*Data{midData},
			Children: []*Node{
				// left,
				node,
				right,
			},
		}

		// left.Parent, right.Parent = newRoot, newRoot
		node.Parent, right.Parent = newRoot, newRoot
		tree.Root = newRoot
		return

	}

	// Okey... there is a parent, could he host it ?

	parent := node.Parent

	if len(parent.Ids) < tree.UpperLimit {
		// Mid index
		insertIndex := sort.Search(len(parent.Ids), func(i int) bool { return midData.Id < parent.Ids[i].Id })

		parent.Ids = append(parent.Ids[:insertIndex], append([]*Data{midData}, parent.Ids[insertIndex:]...)...)
		parent.Children = append(parent.Children[:insertIndex+1], append([]*Node{right}, parent.Children[insertIndex+1:]...)...)

		right.Parent = parent

	} else {
		// Well, it seems whe need to rebalance it too.
		insertIndex := sort.Search(len(parent.Ids), func(i int) bool { return right.Ids[0].Id < parent.Ids[i].Id })
		parent.Children = append(parent.Children[:insertIndex+1], append([]*Node{right}, parent.Children[insertIndex+1:]...)...)
		right.Parent = parent

		log.Println("Parent is full, relabalncing is needed")
		tree.BalanceTree(parent, midData)
	}

}

func (tree *BTree) Insert(node *Node, data *Data) {
	// we are missing
	// to saturate a leaf until it needs to go to the upper level an insert there if the upper level is not the root
	// When splitting how are we going to manage the keys distribuition on the children pointers
	// Check memory for the dist

	// If there is not root, first data inserted
	if node == nil {
		log.Println("We are creating a new root")
		newNode := Node{
			Ids: make([]*Data, 0),
		}
		newNode.Ids = append(newNode.Ids, data)
		tree.Root = &newNode
		return
	}

	// I dont have children, Yuju!
	if len(node.Children) <= 0 {
		log.Println("Inserting on leaf")

		if len(node.Ids) < tree.UpperLimit {
			// There where space!
			node.Ids = append(node.Ids, data)
			sort.Slice(node.Ids,
				func(i, j int) bool {
					return node.Ids[i].Id < node.Ids[j].Id
				},
			)
		} else {
			// There is not space, time to balance
			tree.BalanceTree(node, data)
		}
		return
	}

	// Shit, I have children
	log.Println("Looking for a sibling")
	pointer := len(node.Ids)
	for i, nodeData := range node.Ids {
		if data.Id < nodeData.Id {
			pointer = i
			break
		}
	}
	tree.Insert(node.Children[pointer], data)

}

func NewBTree(order int) *BTree {
	return &BTree{
		Root:       nil,
		UpperLimit: order,
		LowerLimit: order / 2,
	}
}

func main() {
}
