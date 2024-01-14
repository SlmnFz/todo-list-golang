package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// TodoItem represents a single to-do item.
type TodoItem struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// TodoList is a slice of TodoItem.
type TodoList []TodoItem

// AddItem adds a new item to the to-do list.
func (t *TodoList) AddItem(item TodoItem) {
	*t = append(*t, item)
}

// SaveToFile saves the to-do list to a JSON file.
func (t *TodoList) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromFile loads the to-do list from a JSON file.
func (t *TodoList) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, t)
}

// Print prints the to-do list in a human-readable format.
func (t *TodoList) Print() {
	for _, item := range *t {
		status := "pending"
		if item.Completed {
			status = "completed"
		}
		fmt.Printf("Title: %v - Status: %v\n", item.Title, status)
	}
}

// Create a new ToDo by getting user input
func (t *TodoList) createNewToDo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Print("✅ New Item added to the list.\n")
	newItem := TodoItem{Title: title, Completed: false}
	t.AddItem(newItem)
}

// Finish the given Item
func (t *TodoList) finishAnItem(title string) {
	for index, item := range *t {
		if item.Title == title {
			(*t)[index].Completed = true
			break
		}
	}
	err := t.SaveToFile("todo.json")
	if err != nil {
		fmt.Printf("❗️ Error saving to file: %v\n", err)
		return
	}
	fmt.Printf("✅ Item %v has been completed\n", title)
}

func main() {
	// Initialize an empty TodoList
	todoList := &TodoList{}

	// Load existing items from the file
	err := todoList.LoadFromFile("todo.json")
	if err != nil {
		log.Println("Failed to load to-do list:", err)
	}

	todoList.createNewToDo()
	todoList.finishAnItem("new")

	// Save the updated list back to the file
	err = todoList.SaveToFile("todo.json")
	if err != nil {
		log.Fatal("Failed to save to-do list:", err)
	}

	// Print the to-do list in JSON format
	todoList.Print()
}
