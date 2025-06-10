package task

import (
	"cli/internal/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

// UUID creation
func UUID() uuid.UUID {
	Task_id := uuid.Must(uuid.NewRandom())
	// Task_id, err := uuid.Parse("b286d4de-103e-4be0-b99c-72cf005c8a63")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	return Task_id
}

// UUID Parser
func UUIDParse() uuid.UUID {
	UidValue, err := uuid.Parse("b091a34d-2f8b-4d00-8220-f62d540bb4e0")
	if err != nil {
		fmt.Println(err)
	}

	return UidValue
}

//Checks a file exist or not.

func FileCreationCheck(Filename string) {

	if _, err := os.Stat(Filename); os.IsNotExist(err) {
		fmt.Println("File doesn't exist.. Creating...")

		//creating empty slice
		emptyData := []model.Model{}

		data, _ := json.MarshalIndent(emptyData, "", " ")

		fmt.Println("data :", string(data))

		//creating file
		err := ioutil.WriteFile(Filename, data, 0644)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		fmt.Println("Error checking file :", err)
	} else {
		fmt.Println("File already exist !!!")
	}

}

// Read the file data
func ReadFiles(Filename string) []byte {

	//read the data
	read, err := ioutil.ReadFile(Filename)
	if err != nil {
		panic(err)
	}

	//fmt.Println("Reading data from file:", string(read))
	return read
}

// List Task
func ListTask() []model.Model {

	//load task from storage

	file, err := ioutil.ReadFile("./output/dummy.json")
	if err != nil {
		panic(err)
	}

	var list []model.Model

	err = json.Unmarshal(file, &list)
	if err != nil {
		fmt.Println("Error inside unmarshal:", err)
	}
	return list
}

// Appending or inserting new task into file
func AppendTaskIntoFile(Filename string, addtask model.Model) {

	var res []model.Model
	readFile := ReadFiles(Filename)
	err := json.Unmarshal(readFile, &res)
	if err != nil {
		panic(err)
	}

	res = append(res, addtask)

	file, err := json.MarshalIndent(res, "", " ")
	err = ioutil.WriteFile(Filename, file, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Task completed successfully !!!")
}

// Updating task details
func UpdateTask(Filename string, TaskId uuid.UUID, Title string, Status string) {

	list := ListTask()

	updated := false
	for i, j := range list {

		if TaskId == j.ID {

			list[i].Status = Status
			list[i].Title = Title
			//list[i].CreatedAt = time.Now()
			updated = true
			break
		}

	}

	if !updated {
		fmt.Println("TaskId not found")
		return
	}

	file, err := json.MarshalIndent(list, "", " ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(Filename, file, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Task Updated successfully !!!")

}

// Delete specific task
func DeleteTask(Filename string, taskId uuid.UUID) {

	data := ListTask()
	TestValue := false
	var updatedValue []model.Model

	for _, j := range data {
		if j.ID != taskId {
			updatedValue = append(updatedValue, j)
		} else {
			TestValue = true
		}

	}

	if !TestValue {

		fmt.Println("Task Id not found")
		return
	}

	val, err := json.MarshalIndent(updatedValue, "", " ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(Filename, val, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TaskId: %+v\n", taskId)
	fmt.Println("Task Deleted Successfully !!!")
}

// Delete All existing Task
func DeleteAllTask(Filename string) {

	dats := []model.Model{}

	val, err := json.MarshalIndent(dats, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("val: %+v", string(val))
	err = ioutil.WriteFile(Filename, val, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted All Task !!!")
}
