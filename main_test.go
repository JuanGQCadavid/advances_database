package main

import (
	"log"
	"testing"
)

func TestFirstNode(t *testing.T) {
	bTree := NewBTree(5)
	bTree.Insert(bTree.Root, &Data{
		Id:       1,
		Position: 0,
	})

	if bTree.Root != nil {
		if len(bTree.Root.Ids) == 1 {
			for i := range bTree.Root.Ids {
				log.Println(bTree.Root.Ids[i].Id)
			}

		} else {
			log.Panic("We fuck up!")
		}
	} else {
		log.Panic("We fuck up twice!")
	}

}

func TestFirstNodeMoreData(t *testing.T) {
	bTree := NewBTree(5)
	bTree.Insert(bTree.Root, &Data{
		Id:       4,
		Position: 0,
	})
	bTree.Insert(bTree.Root, &Data{
		Id:       2,
		Position: 0,
	})
	bTree.Insert(bTree.Root, &Data{
		Id:       5,
		Position: 0,
	})
	bTree.Insert(bTree.Root, &Data{
		Id:       8,
		Position: 0,
	})

	if bTree.Root != nil {
		if len(bTree.Root.Ids) > 0 {
			for i := range bTree.Root.Ids {
				log.Println(bTree.Root.Ids[i].Id)
			}

		} else {
			log.Panic("We fuck up!")
		}
	} else {
		log.Panic("We fuck up twice!")
	}

}

func TestFirstNodeMoreDatav2(t *testing.T) {
	bTree := NewBTree(5)
	log.Println("4")
	bTree.Insert(bTree.Root, &Data{
		Id:       4,
		Position: 0,
	})

	log.Println("2")
	bTree.Insert(bTree.Root, &Data{
		Id:       2,
		Position: 0,
	})
	log.Println("5")
	bTree.Insert(bTree.Root, &Data{
		Id:       5,
		Position: 0,
	})

	log.Println("8")
	bTree.Insert(bTree.Root, &Data{
		Id:       8,
		Position: 0,
	})

	log.Println("3")
	bTree.Insert(bTree.Root, &Data{
		Id:       3,
		Position: 0,
	})

	log.Println("9")
	bTree.Insert(bTree.Root, &Data{
		Id:       9,
		Position: 0,
	})

	log.Println("10")
	bTree.Insert(bTree.Root, &Data{
		Id:       10,
		Position: 0,
	})

	log.Println("1")
	bTree.Insert(bTree.Root, &Data{
		Id:       1,
		Position: 0,
	})

	log.Println("-1")
	bTree.Insert(bTree.Root, &Data{
		Id:       -1,
		Position: 0,
	})
	if bTree.Root != nil {
		if len(bTree.Root.Ids) > 0 {
			for i := range bTree.Root.Ids {
				log.Println(bTree.Root.Ids[i].Id)
			}

		} else {
			log.Panic("We fuck up!")
		}
	} else {
		log.Panic("We fuck up twice!")
	}

}

func TestFirstNodeMoreDatav3(t *testing.T) {
	bTree := NewBTree(3)
	log.Println("4")
	bTree.Insert(bTree.Root, &Data{
		Id:       4,
		Position: 0,
	})

	log.Println("2")
	bTree.Insert(bTree.Root, &Data{
		Id:       2,
		Position: 0,
	})
	log.Println("5")
	bTree.Insert(bTree.Root, &Data{
		Id:       5,
		Position: 0,
	})

	log.Println("8")
	bTree.Insert(bTree.Root, &Data{
		Id:       8,
		Position: 0,
	})

	log.Println("3")
	bTree.Insert(bTree.Root, &Data{
		Id:       3,
		Position: 0,
	})

	log.Println("9")
	bTree.Insert(bTree.Root, &Data{
		Id:       9,
		Position: 0,
	})

	log.Println("10")
	bTree.Insert(bTree.Root, &Data{
		Id:       10,
		Position: 0,
	})

	log.Println("1")
	bTree.Insert(bTree.Root, &Data{
		Id:       1,
		Position: 0,
	})

	log.Println("-1")
	bTree.Insert(bTree.Root, &Data{
		Id:       -1,
		Position: 0,
	})

	bTree.Insert(bTree.Root, &Data{
		Id:       11,
		Position: 0,
	})

	bTree.Insert(bTree.Root, &Data{
		Id:       12,
		Position: 0,
	})

	bTree.Insert(bTree.Root, &Data{
		Id:       13,
		Position: 0,
	})

	bTree.Insert(bTree.Root, &Data{
		Id:       7,
		Position: 0,
	})

	bTree.Insert(bTree.Root, &Data{
		Id:       6,
		Position: 0,
	})

	bTree.Insert(bTree.Root, &Data{
		Id:       14,
		Position: 0,
	})

	bTree.Insert(bTree.Root, &Data{
		Id:       15,
		Position: 0,
	})
	bTree.Insert(bTree.Root, &Data{
		Id:       16,
		Position: 0,
	})

	bTree.Insert(bTree.Root, &Data{
		Id:       17,
		Position: 0,
	})
	if bTree.Root != nil {
		if len(bTree.Root.Ids) > 0 {
			for i := range bTree.Root.Ids {
				log.Println(bTree.Root.Ids[i].Id)
			}

		} else {
			log.Panic("We fuck up!")
		}
	} else {
		log.Panic("We fuck up twice!")
	}

}
