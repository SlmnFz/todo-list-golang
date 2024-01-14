package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"strings"
)

var (
	indexColor  = color.New(color.FgHiYellow).SprintfFunc()
	titleColor  = color.New(color.FgCyan).SprintfFunc()
	pendingColor = color.New(color.FgRed).SprintfFunc()
	completedColor = color.New(color.FgGreen).SprintfFunc()
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
	for index, item := range *t {
		status := pendingColor("pending")
		if item.Completed {
			status = completedColor("completed")
		}
		var emoji string
		if index%2 == 0 {
			emoji = "⚫️"
		} else {
			emoji = "⚪️"
		}
		fmt.Printf("%v %v. Title: %v - Status: %v\n", emoji, indexColor("%d", index+1), titleColor(item.Title), status)
	}
}

// Create a new ToDo given by the user
func (t *TodoList) createNewToDo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Print("✅ New Item added to the list.\n")
	newItem := TodoItem{Title: title, Completed: false}
	t.AddItem(newItem)
}

// Finish the given Item by the user
func (t *TodoList) finishAnItem() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the title you want completed: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	for index, item := range *t {
		if item.Title == title {
			(*t)[index].Completed = true
			break
		}
		if index+1 == len(*t) && item.Title != title {
			fmt.Printf("❌ Item %v Was Not Found.\n", title)
			return
		}
	}
	fmt.Printf("✅ Item %v has been completed\n", title)
}

// Delete the given Item by the user
func (t *TodoList) deleteAnItem() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the title you want deleted: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	for index, item := range *t {
		if item.Title == title {
			*t = append((*t)[:index], (*t)[index+1:]...)
			break
		}
		if index+1 == len(*t) && item.Title != title {
			fmt.Printf("❌ Item %v Was Not Found.\n", title)
			return
		}
	}
	fmt.Printf("✅ Item %v has been deleted.\n", title)
}

// Edit an Item title
func (t *TodoList) editAnItem() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the title you want to edit: ")
	title, _ := reader.ReadString('\n')
	newTitle := title
	title = strings.TrimSpace(title)
	for index, item := range *t {
		if item.Title == title {
			fmt.Print("Enter the new title: ")
			newTitle, _ = reader.ReadString('\n')
			newTitle = strings.TrimSpace(newTitle)
			(*t)[index].Title = newTitle
			break
		}
		if index+1 == len(*t) && item.Title != title {
			fmt.Printf("❌ Item %v Was Not Found.\n", title)
			return
		}
	}
	fmt.Printf("✅ Item %v has been edited ==> New Title: %v\n", title, newTitle)
}

// Show ToDo App Menu
func (t *TodoList) menu() {
	fmt.Println("\n📍📍📍 ToDo App 📍📍📍")
	fmt.Println("1. Create new ToDo")
	fmt.Println("2. Finish A ToDo")
	fmt.Println("3. Edit A ToDo")
	fmt.Println("4. Delete A ToDo")
	fmt.Println("5. Save")
	fmt.Println("6. Show My ToDo List")
	fmt.Println("7. Exit")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your choice: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		t.createNewToDo()
	case "2":
		t.finishAnItem()
	case "3":
		t.editAnItem()
	case "4":
		t.deleteAnItem()
	case "5":
		err := t.SaveToFile("todo.json")
		if err != nil {
			log.Println("Failed to save to-do list:", err)
		} else {
			fmt.Println("To-Do list saved.")
		}
	case "6":
		t.Print()
	case "7":
		fmt.Println("👋 Bye Bye 👋")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice.")
	}

	// Recursively call the menu function to show the menu again
	t.menu()
}

func main() {
	// Initialize an empty TodoList
	todoList := &TodoList{}

	// Load existing items from the file
	err := todoList.LoadFromFile("todo.json")
	if err != nil {
		log.Println("Failed to load to-do list:", err)
	}

	todoList.menu()
}
