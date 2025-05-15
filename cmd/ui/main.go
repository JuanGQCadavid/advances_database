package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/advances_database/core"
	"github.com/manifoldco/promptui"
)

const (
	Insert string = "Insert"
	Search string = "Search"
	Exit   string = "Exit"
)

var (
	bTree *core.BTree
	DB    []*core.Student
)

func init() {
	bTree = core.NewBTree(5)
	DB = make([]*core.Student, 0)

}

func main() {
	for {
		clearScreen()
		prompt := promptui.Select{
			Label: "Select operation",
			Items: []string{Insert, Search, Exit},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result)

		switch result {
		case Insert:
			insert()
		case Search:
			search()
		case Exit:
			return
		}
	}

}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func search() {
	number := askString("Number")
	id, err := strconv.ParseInt(number, 10, 64)

	if err != nil {
		fmt.Println("The id is not valid, cast wrong", err.Error())
		return
	}

	foundData := bTree.Search(bTree.Root, int(id))

	if foundData != nil {
		fmt.Printf("Data index: %d,%d \n", foundData.Id, foundData.Position)
		fmt.Println("Database data: ", DB[foundData.Position].ToString())

	} else {
		fmt.Println("No data found on index")
	}

	askString("Press enter to continue")

}

func insert() {

	number := askString("Number")
	name := askString("name")
	gender := askString("gender")
	age := askString("age")
	city := askString("city")

	id, err := strconv.ParseInt(number, 10, 64)

	if err != nil {
		fmt.Println("The id is not valid, cast wrong", err.Error())
		return
	}

	newStudent := &core.Student{
		Number: int(id),
		Name:   name,
		Gender: gender,
		Age:    age,
		City:   city,
	}

	fmt.Println("You Insert: ", newStudent.ToString())

	DB = append(DB, newStudent)
	bTree.Insert(bTree.Root, &core.Data{
		Id:       int(id),
		Position: len(DB) - 1,
	})

}

func askString(label string) string {
	// validate := func(input string) error {
	// 	_, err := strconv.ParseFloat(input, 64)
	// 	if err != nil {
	// 		return errors.New("Invalid number")
	// 	}
	// 	return nil
	// }

	prompt := promptui.Prompt{
		Label: label,
		// Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result

	// fmt.Printf("You choose %q\n", result)
}
