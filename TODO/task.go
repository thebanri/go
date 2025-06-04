package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"desc"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdateAt    string `json:"updatedAt"`
}

var StatusData = []string{"todo", "done", "in-progress"}

var fileName string = "/etc/taskcli/tasks.json"

func fileExist() {

	if _, err := os.Stat(fileName); err == nil {
		return
	} else if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error in file creating: ", err)
			return
		}
		defer file.Close()
	} else {
		fmt.Println("Error in file cheking: ", err)
		return
	}

}

func ReadTasks() []Task {

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error in file reading:", err)
	}
	var tasks []Task
	json.Unmarshal(fileContent, &tasks)

	return tasks

}

func WriteTasks(task []Task) {

	file, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		fmt.Println("Error in transforming to JSON data:", err)
	}

	err = os.WriteFile(fileName, file, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

}

func GenerateId() int {

	tasks := ReadTasks()
	if len(tasks) == 0 {
		return 1
	}

	maxId := 0

	for _, task := range tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}
	return maxId + 1

}

func UpdateTodo(id int, desc string) {
	tasks := ReadTasks()
	found := false
	for i := range tasks {
		if tasks[i].Id == id {
			tasks[i].Description = desc
			tasks[i].UpdateAt = time.Now().Format("Mon 01-02-2006")
			found = true
			break
		}
	}

	if found {
		WriteTasks(tasks)
		fmt.Println("Task description updated!", "\nTask ID:", id)
	} else {
		fmt.Println("Task ID not found!")
	}

}

func UpdateStatus(statusCode, id int) {
	tasks := ReadTasks()
	for i := range tasks {
		if tasks[i].Id == id {
			tasks[i].Status = StatusData[statusCode]
			tasks[i].UpdateAt = time.Now().Format("Mon 01-02-2006")
			break
		}
	}
	WriteTasks(tasks)

	fmt.Println("Task Status Updated! ", "\nTask ID:", id)
}

func NewTodo(desc string) {

	id := GenerateId()
	createdAt := time.Now().Format("Mon 01-02-2006")

	updatedAt := time.Now().Format("Mon 01-02-2006")

	newTodo := &Task{Id: id, Description: desc, Status: StatusData[0], CreatedAt: createdAt, UpdateAt: updatedAt}

	tasks := ReadTasks()

	tasks = append(tasks, *newTodo)

	WriteTasks(tasks)

	fmt.Printf("New Task Created: %+v\n", *newTodo)

}

func ListDone() {
	tasks := ReadTasks()
	doneTasks := []Task{}

	for _, task := range tasks {
		if task.Status == "done" {
			doneTasks = append(doneTasks, task)
		}
	}

	if len(doneTasks) == 0 {
		fmt.Println("No tasks marked as done.")
		return
	}

	for _, task := range doneTasks {
		fmt.Println("ID:", task.Id, "\nDescription:", task.Description, "\nCreated Date:", task.CreatedAt, "\nLast Updated Date:", task.UpdateAt, "\nStatus:", task.Status)
	}
}

func ListInProgress() {
	tasks := ReadTasks()
	inProgressTasks := []Task{}

	for _, task := range tasks {
		if task.Status == "in-progress" {
			inProgressTasks = append(inProgressTasks, task)
		}
	}

	if len(inProgressTasks) == 0 {
		fmt.Println("No tasks marked as in-progress.")
		return
	}

	for _, task := range inProgressTasks {
		fmt.Println("ID:", task.Id, "\nDescription:", task.Description, "\nCreated Date:", task.CreatedAt, "\nLast Updated Date:", task.UpdateAt, "\nStatus:", task.Status)
	}
}

func ListTodo() {
	tasks := ReadTasks()
	todoTasks := []Task{}

	for _, task := range tasks {
		if task.Status == "todo" {
			todoTasks = append(todoTasks, task)
		}
	}

	if len(todoTasks) == 0 {
		fmt.Println("No tasks to do.")
		return
	}

	for _, task := range todoTasks {
		fmt.Println("ID:", task.Id, "\nDescription:", task.Description, "\nCreated Date:", task.CreatedAt, "\nLast Updated Date:", task.UpdateAt, "\nStatus:", task.Status)
	}
}

func DeleteTodo(id int) {

	tasks := ReadTasks()
	updatedTasks := []Task{}

	for _, task := range tasks {
		if task.Id == id {

		} else {
			updatedTasks = append(updatedTasks, task)
		}
	}

	WriteTasks(updatedTasks)
}

func ListAll() {
	tasks := ReadTasks()

	for _, task := range tasks {

		fmt.Println("ID:", task.Id, "\nDescription:", task.Description, "\nCreated Date:", task.CreatedAt, "\nLast Updated Date:", task.UpdateAt, "\nStatus:", task.Status)

	}
}

func main() {
	fileExist()

	args := os.Args[1:]
	argLength := len(args)

	if argLength == 0 {
		fmt.Println("Please provide a command. Use 'task help' for usage.")
		return
	}

	command := args[0]

	switch command {
	case "list":
		if argLength == 1 {
			ListAll()
		} else if argLength == 2 {
			switch args[1] {
			case "todo":
				ListTodo()
			case "done":
				ListDone()
			case "in-progress":
				ListInProgress()
			default:
				fmt.Printf("Unknown list type: '%s'. Use 'task help' for usage.\n", args[1])
				return
			}
		} else {
			fmt.Println("Usage: task list [todo|done|in-progress]")
			return
		}
	case "add":
		if argLength < 2 {
			fmt.Println("Usage: task add \"<description>\"")
			return
		}
		NewTodo(args[1])
	case "mark-in-progress":
		if argLength < 2 {
			fmt.Println("Usage: task mark-in-progress <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID. Please provide a number.")
			return
		}
		UpdateStatus(2, id)
	case "mark-done":
		if argLength < 2 {
			fmt.Println("Usage: task mark-done <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID. Please provide a number.")
			return
		}
		UpdateStatus(1, id)
	case "update":
		if argLength < 3 {
			fmt.Println("Usage: task update <id> \"<new description>\"")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID. Please provide a number.")
			return
		}
		UpdateTodo(id, args[2])
	case "delete":
		if argLength < 2 {
			fmt.Println("Usage: task delete <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID. Please provide a number.")
			return
		}
		DeleteTodo(id)
	case "help":
		fmt.Println("Commands:")
		fmt.Println("  list                           - List all tasks")
		fmt.Println("  add \"<description>\"          - Add a new task")
		fmt.Println("  update <id> \"<new description>\" - Update a task's description")
		fmt.Println("  mark-in-progress <id>        - Mark a task as 'in-progress'")
		fmt.Println("  mark-done <id>               - Mark a task as 'done'")
		fmt.Println("  delete <id>                  - Delete a task")
		fmt.Println("  help                           - Show this help message")
	default:
		fmt.Printf("Unknown command: '%s'. Use 'task help' for usage.\n", command)
	}
}
