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
	if node.Children == nil {
		log.Println("Inserting on leaf")
		node.Ids = append(node.Ids, data)
		sort.Slice(node.Ids,
			func(i, j int) bool {
				return node.Ids[i].Id < node.Ids[j].Id
			},
		)

		if len(node.Ids) >= tree.UpperLimit {
			log.Println("We are creating new children")
			midIndex := len(node.Ids) / 2
			midData := node.Ids[midIndex]
			left := &Node{Ids: append([]*Data{}, node.Ids[:midIndex]...)}
			right := &Node{Ids: append([]*Data{}, node.Ids[midIndex+1:]...)}

			if node.Parent == nil {
				newRoot := Node{
					Ids: []*Data{midData},
				}
				newRoot.Children = make([]*Node, 0)
				newRoot.Children = append(newRoot.Children, left)
				newRoot.Children = append(newRoot.Children, right)

				left.Parent = &newRoot
				right.Parent = &newRoot

				tree.Root = &newRoot
			}

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
