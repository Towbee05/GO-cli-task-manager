package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type TaskStoreStruct struct {
	Id          int
	Name        string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var TaskId uint = 0
var TaskStore = make([]TaskStoreStruct, 0)

func AddTask(task string, descriptions ...string) string {
	ReadJsonFile()
	var description string = ""
	if len(descriptions) > 0 {
		description = descriptions[0]
	}

	var userTask = TaskStoreStruct{
		Id:          len(TaskStore) + 1,
		Name:        task,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	TaskId++
	TaskStore = append(TaskStore, userTask)
	data, err := json.MarshalIndent(TaskStore, "", "  ")
	if err != nil {
		fmt.Printf("An error occured while converting into json file %v \n", err)
	}
	err = os.WriteFile("data.json", data, 0644)
	if err != nil {
		fmt.Printf("An error occured while writing into json file %v \n", err)
	}

	return fmt.Sprintf("Task added successfully (ID: %v)", userTask.Id)
}

func ReadJsonFile() {
	var fileContent, readErr = os.ReadFile("data.json")
	if readErr != nil {
		fmt.Printf("An error occured while reading json file %v \n", readErr)
	}
	if len(fileContent) == 0 {
		return
	}
	err := json.Unmarshal(fileContent, &TaskStore)
	if err != nil {
		fmt.Printf("An error occured while converting json file %v \n", readErr)
	}

}

// Listing all tasks
func List(filters ...string) []string {
	ReadJsonFile()
	var tasks []string = make([]string, 0)
	var filter string = ""
	if len(filter) > 0 {
		filter = filters[0]
	}

	fmt.Println(filter)
	switch filter {
	case "done":
		tasks = handleFilter("done", tasks)
		return tasks
	case "todo":
		tasks = handleFilter("todo", tasks)
		return tasks
	case "in-progress":
		tasks = handleFilter("in-progress", tasks)
		return tasks
	default:
		for _, task := range TaskStore {
			tasks = append(tasks, task.Name)
		}
		return tasks
	}
}

func handleFilter(filter string, tasks []string) []string {
	for _, task := range TaskStore {
		if task.Status == filter {
			tasks = append(tasks, task.Name)
		}
	}
	return tasks
}

func UpdateTask(id int, name string) string {
	ReadJsonFile()
	if id < 0 {
		return fmt.Sprintln("Provided ID should be bigger than zero")
	}
	var index int = id - 1
	if index > len(TaskStore)-1 {
		return fmt.Sprintf("Could not access task with (ID: %v \n)", id)
	}
	var task TaskStoreStruct = TaskStore[index]
	task.Name = name
	task.UpdatedAt = time.Now()
	return fmt.Sprintf("Task updated successfully (ID: %v)", index)
}

// func DeleteTask(id int) string {
// 	ReadJsonFile()
// 	if id < 0 {
// 		return fmt.Sprintln("Provided ID should be bigger than zero")
// 	}
// 	var index int = id - 1
// 	if index > len(TaskStore)-1 {
// 		return fmt.Sprintf("Could not access task with (ID: %v \n)", id)
// 	}
// 	var task TaskStoreStruct = TaskStore[index]

// 	return fmt.Sprintf("Task deleted successfully (ID: %v)", index)
// }
