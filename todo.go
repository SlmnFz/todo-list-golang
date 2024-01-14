package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

func main() {
	// Initialize an empty TodoList
	todoList := &TodoList{}

	// Load existing items from the file
	err := todoList.LoadFromFile("todo.json")
	if err != nil {
		log.Println("Failed to load to-do list:", err)
	}

	// Example of adding a new item to the list
	newItem := TodoItem{Title: "Learn Go v3", Completed: false}
	todoList.AddItem(newItem)

	// Save the updated list back to the file
	err = todoList.SaveToFile("todo.json")
	if err != nil {
		log.Fatal("Failed to save to-do list:", err)
	}

	// Print the to-do list in JSON format
	todoList.Print()
}
