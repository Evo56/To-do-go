package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type todo struct {
	id        int
	Title     string
	Completed bool
}

var todos []todo
var nextId = 1

func addTodo(title string) {
	todo := todo{
		id:        nextId,
		Title:     title,
		Completed: false,
	}

	todos = append(todos, todo)
	nextId++

	fmt.Println("Todo added successfully")
}

func listTodos() {
	if len(todos) == 0 {
		fmt.Println("No todos Found")
		return
	}

	fmt.Println("Your Todos")

	for _, todo := range todos {
		status := "not completed"

		if todo.Completed {
			status = "completed"
		}

		fmt.Printf("%v. %s %s\n", todo.id, todo.Title, status)
	}
}

func completedtodo(id int) {
	for i := range todos {
		if todos[i].id == id {
			todos[i].Completed = true
			fmt.Println("todo completed!")
			return
		}
	}

	fmt.Println("todo not found")
}

func deletetodo(id int) {
	newTodos := []todo{}

	for _, todo := range todos {
		if todo.id != id {
			newTodos = append(newTodos, todo)
		}
	}

	todos = newTodos

	fmt.Println("Todo deleted!")
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("1 Add Todo")
		fmt.Println("2 List Todos")
		fmt.Println("3 Complete Todo")
		fmt.Println("4 Delete Todo")
		fmt.Println("5 Exit")

		fmt.Print("choose an option:")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("please enter a number")
			continue
		}
		switch choice {

		case 1:
			fmt.Print("Enter todo:")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			if title == "" {
				fmt.Println("Todo cannot be empty.")
				continue
			}

			addTodo(title)

		case 2:
			listTodos()

		case 3:
			fmt.Print("Enter todo ID:")
			input, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid Id")
				continue
			}
			completedtodo(id)

		case 4:
			fmt.Print("Enter todo ID: ")

			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			id, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("Invalid ID.")
				continue
			}

			deletetodo(id)

		case 5:
			fmt.Println("Thank you!")
			return

		default:
			fmt.Println("Invalid option.")
		}

	}

}
