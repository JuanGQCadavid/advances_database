package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/advances_database/core"
)

const (
	CSV_PATH = "./files/stud_record.csv"
)

type PairId struct {
	Id       int
	Position int
}

var (
	bTree *core.BTree
	DB    []*core.Student
	Index []*PairId
)

func init() {
	bTree = core.NewBTree(5)
	DB = make([]*core.Student, 0)
	Index = make([]*PairId, 0)

}

func getRecords() [][]string {
	file, err := os.Open(CSV_PATH)
	if err != nil {
		log.Fatal("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV:", err)
		return nil
	}

	return records
}

func main() {

	times := 10
	targetToSearch := 200

	records := getRecords()[1:]
	timesPerIndex := map[string][]time.Duration{
		"bTree": make([]time.Duration, times),
		"link":  make([]time.Duration, times),
	}

	searchPerIndex := map[string][]time.Duration{
		"bTree": make([]time.Duration, times),
		"link":  make([]time.Duration, times),
	}

	for indexType := range timesPerIndex {
		log.Println("----------")
		log.Println("indexType: ", indexType)
		for i := range timesPerIndex[indexType] {
			bTree = core.NewBTree(5)
			DB = make([]*core.Student, 0)
			Index = make([]*PairId, 0)

			if indexType == "bTree" {
				timesPerIndex[indexType][i] = test(records, true, false)
				searchPerIndex[indexType][i] = testSearch(targetToSearch, true, false)
			}
			if indexType == "link" {
				timesPerIndex[indexType][i] = test(records, false, true)
				searchPerIndex[indexType][i] = testSearch(targetToSearch, false, true)
			}

		}
		log.Println("Insert: ", indexType, timesPerIndex[indexType])
		log.Println("Search: ", indexType, searchPerIndex[indexType])
		log.Println()
	}

}

func testSearch(id int, btree bool, linked bool) time.Duration {

	startTime := time.Now()

	if btree {
		foundData := bTree.Search(bTree.Root, id)
		if foundData == nil || foundData.Id != id {
			log.Panic("Shit, the data was nor found")
		}

	}

	if linked {
		found := false
		for _, p := range Index {
			if p.Id == id {
				found = true
				break
			}
		}

		if !found {
			log.Panic("Shit, the data was nor found")
		}
	}

	return time.Since(startTime)
}

func test(records [][]string, btree bool, linked bool) time.Duration {

	startTime := time.Now()
	for _, data := range records {

		id, err := strconv.ParseInt(data[0], 10, 64)

		if err != nil {
			log.Fatalln("The id is not valid, cast wrong", err.Error())
		}

		newStudent := &core.Student{
			Number: int(id),
			Name:   data[1],
			Gender: data[2],
			Age:    data[3],
			City:   data[4],
		}

		DB = append(DB, newStudent)

		if btree {
			bTree.Insert(bTree.Root, &core.Data{
				Id:       int(id),
				Position: len(DB) - 1,
			})
		}

		if linked {
			Index = append(Index, &PairId{
				Id:       int(id),
				Position: len(DB) - 1,
			})
		}
	}

	delta := time.Since(startTime)
	log.Println("it tooks: ", delta)
	return delta

}
