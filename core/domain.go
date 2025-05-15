package core

import "fmt"

type Student struct {
	Number int
	Name   string
	Gender string
	Age    string
	City   string
}

func (std *Student) ToString() string {
	return fmt.Sprintf("\nId:%d,\nName:%s,\nGender:%s,\nAge:%s,\nCity:%s\n", std.Number, std.Name, std.Gender, std.Age, std.City)
}
