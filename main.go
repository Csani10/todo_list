package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
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

	if strings.ToLower(input) == "y\n" {
		fmt.Println("Type 'quit' to stop")
		for {
			fmt.Print("Task: ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			if strings.ToLower(input) == "quit\n" {
				break
			}
			task_str := input

			date := time.Now()
			for {
				fmt.Print("Due date (format: '2006-06-01'): ")
				input, err = reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error:", err)
					continue
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
	} else {
		fmt.Print("Task: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		task_str := input

		date := time.Now()
		for {
			fmt.Print("Due date (format: '2006-06-01'): ")
			input, err = reader.ReadString('\n')
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

	fmt.Println("Tasks created!")
}

func List() {
	fmt.Println("Listing tasks")

	file, err := os.Open(os.Getenv("HOME") + "/.config/todo_list/tasks")
	if err != nil {
		fmt.Println("Couldnt open tasks file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line_b, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		line := string(line_b[:])

		task := tasks.Deserialize(line)
		fmt.Println(task.GetTask())
	}
}

func CheckConfigDir() {
	_, err := os.Stat("~/.config/todo_list/")
	if err == nil {
		return
	}
	if errors.Is(err, fs.ErrNotExist) {
		os.Mkdir("~/.config/todo_list/", os.ModeDir)
	}

	os.Create("~/.config/todo_list/tasks")
}

func main() {
	if len(os.Args) < 2 {
		Usage()
		return
	}

	switch os.Args[1] {
	case "create":
		CheckConfigDir()
		Create()
	case "list":
		CheckConfigDir()
		List()
	default:
		Usage()
		fmt.Println("Error:", os.Args[1], "operation doesnt exist")
	}
}
