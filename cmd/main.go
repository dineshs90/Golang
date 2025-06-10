package main

import (
	"cli/internal/controller"
	"cli/internal/task"
	"fmt"
)

func main() {

	fmt.Println("CLI Project")

	Filename := "./output/dummy.json"

	//checking file creation
	task.FileCreationCheck(Filename)

	controller.CLI(Filename)

}
