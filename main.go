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
			left, middle, right := node.Ids[0:tree.UpperLimit/2], node.Ids[tree.UpperLimit/2], node.Ids[(tree.UpperLimit/2)+1:]
			nodeLeft := Node{
				Ids: left,
			}
			nodeRight := Node{
				Ids: right,
			}

			if node.Parent == nil {
				newRoot := Node{
					Ids: []*Data{middle},
				}
				newRoot.Children = make([]*Node, 0)
				newRoot.Children = append(newRoot.Children, &nodeLeft)
				newRoot.Children = append(newRoot.Children, &nodeRight)

				nodeLeft.Parent = &newRoot
				nodeRight.Parent = &newRoot

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
