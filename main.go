package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"todo_list/tasks"
)

func Usage() {
	fmt.Println("Usage: ./todo_list [operation]")
	fmt.Println("Operations:")
	fmt.Println("    create - Create a new task")
	fmt.Println("    list - List tasks")
}

func Create() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Add multiple tasks? (y/N): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	task_array := []tasks.Task{}

	if strings.ToLower(input) == "y" {
		fmt.Printf("Type 'quit' to stop")
		for {
			fmt.Print("Task: ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			if strings.ToLower(input) == "quit" {
				break
			}
			task_str := input

			fmt.Print("Due date (format: '2006-06-01'): ")
			input, err = reader.ReadString('\n')
			date := time.Now()
			for {
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				date, err = time.Parse("2006-06-01", input)
				if err != nil {
					fmt.Println("Invalid format")
					continue
				}
				break
			}

			task := tasks.NewTask(task_str, date)
			task_array = append(task_array, *task)
		}
	} else {
		fmt.Print("Task: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		task_str := input

		fmt.Print("Due date (format: '2006-06-01'): ")
		input, err = reader.ReadString('\n')
		date := time.Now()
		for {
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			date, err = time.Parse("2006-06-01", strings.Trim(input, "\n\x0a"))
			if err != nil {
				fmt.Println("Invalid format")
				continue
			}
			break
		}

		task := tasks.NewTask(task_str, date)
		task_array = append(task_array, *task)
	}

	for taskk := range task_array {
		fmt.Println(taskk)
	}
}

func List() {
	fmt.Println("list")
}

func main() {
	if len(os.Args) < 2 {
		Usage()
		return
	}

	switch os.Args[1] {
	case "create":
		Create()
	case "list":
		List()
	default:
		Usage()
		fmt.Println("Error:", os.Args[1], "operation doesnt exist")
	}
}
