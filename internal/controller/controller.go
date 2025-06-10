package controller

import (
	"bufio"
	"cli/internal/model"
	"cli/internal/task"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func CLIInput() (string, string) {

	//var TaskId string
	var Title string
	var Status string
	//var CreatedAt string
	// fmt.Print("Enter the Task ID :")

	// _, err := fmt.Scan(&TaskId)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Print("Enter the Title :")
	_, err := fmt.Scan(&Title)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Enter the Status :")
	_, err = fmt.Scan(&Status)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Print("Enter the CreatedAt :")
	// _, err = fmt.Scan(&CreatedAt)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//fmt.Println("TaskId: ", TaskId)
	fmt.Println("Title: ", Title)
	fmt.Println("Status: ", Status)
	//fmt.Println("CreatedAt: ", CreatedAt)

	return Title, Status
}

func GetTaskID() uuid.UUID {
	fmt.Print("Enter the Task ID :")

	var TaskId uuid.UUID

	//to get full user input
	Reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the Task ID :")

	input, err := Reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	//to remove newlines or spaces from user input
	input = strings.TrimSpace(input)

	//convert string to UUID type

	TaskId, err = uuid.Parse(input)
	if err != nil {
		fmt.Println(err)
	}

	return TaskId
}

func CLIUpdate() (uuid.UUID, string, string) {

	TaskId := GetTaskID()

	// _, err := fmt.Scan(&TaskId)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	Title, Status := CLIInput()

	return TaskId, Title, Status
}

// func CLIDelete() {

// 	TaskId := GetTaskID()

// }

func CLI(Filename string) {

	var input string
	fmt.Println("**********")
	fmt.Println("1) New Task")
	fmt.Println("2) Update Task")
	fmt.Println("3) Delete Task")
	fmt.Println("4) Delete All Task")
	fmt.Println("5) List All Task")

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
	}

	inp, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}

	if inp == 1 {
		//Sending input to AppendingTaskIntoFile function
		fmt.Println("Creating New Task")
		Title, Completed := CLIInput()
		taskouput := model.Model{ID: task.UUID(), Title: Title, Status: Completed, CreatedAt: time.Now()}

		//Adding into file
		task.AppendTaskIntoFile(Filename, taskouput)
		fmt.Println("Task list after creations: ", string(task.ReadFiles(Filename)))
	} else if inp == 2 {

		fmt.Println("Updating existing Task")
		//Update task details
		TId, Tit, Stat := CLIUpdate()
		task.UpdateTask(Filename, TId, Tit, Stat)

		fmt.Println("Task list after Update: ", string(task.ReadFiles(Filename)))
	} else if inp == 3 {
		//Deleting the task

		task.DeleteTask(Filename, GetTaskID())

		fmt.Println("Task list after deletion: ", string(task.ReadFiles(Filename)))
	} else if inp == 4 {
		//Delete All Task

		task.DeleteAllTask(Filename)
	} else if inp == 5 {
		//List All task

		fmt.Printf(string(task.ReadFiles(Filename)))
	} else {
		fmt.Println("Please enter valid input")
		CLI(Filename)
	}
}
