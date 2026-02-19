package main

import (
	"fmt"

	"github.com/amenshenin/go-leet-code/tests"
)

func main() {
	command := ""
	grade := ""
	task := ""
	fmt.Print("Введите команду: ")
	fmt.Scanln(&command, &grade, &task)
	switch command {
	case "Run":
		error := tests.Run(grade, task)
		if error != nil {
			fmt.Printf("error %s", error.Error())
			fmt.Println("")
		}
	case "GetNewName":
		fmt.Printf("New task name: %s", tests.GetNewName(5))
		fmt.Println("")
	default:
		fmt.Printf("wtf")
		fmt.Println("")
	}
}
