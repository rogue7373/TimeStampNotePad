package main

import (
	"fmt"
	"os"
	"time"
)

const (
	pacificTZ = "America/Los_Angeles"
	file      = "notes.txt"
)

func main() {
	// Create a new file to store the notes.
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Append a new line to the file with the current time stamp.
	addNote := func(note string) {
		t := getTime()
		f.WriteString(fmt.Sprintf("%s: %s\n", t, note))
	}

	// Read the contents of the file and display them.
	displayNotes := func() {
		b, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
	}

	// Usage example
	addNote("This is my first note.")
	addNote("This is my second note.")
	displayNotes()
}

// Get the current time in Pacific Time Zone.
func getTime() string {
	loc, err := time.LoadLocation(pacificTZ)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return time.Now().In(loc).Format("2006-01-02 15:04:05 MST")
}
