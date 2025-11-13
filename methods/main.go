package main

import (
	"fmt"
	"strings"
)

type Notebook struct {
	Title  string
	Notes  []string
	Closed bool
}

func (n Notebook) Count() int {
	return len(n.Notes)
}

func (n *Notebook) Add(note string) {
	if n.Closed {
		fmt.Println("notebook is closed; cannot add note")
		return
	}
	n.Notes = append(n.Notes, note)
}

func (n *Notebook) Close() {
	n.Closed = true
}

func (n Notebook) Summary() string {
	return fmt.Sprintf("%s: %s", n.Title, strings.Join(n.Notes, ", "))
}

func main() {
	nb := &Notebook{Title: "Go Webinar"}

	nb.Add("functions return multiple values")
	nb.Add("struct embedding promotes fields")

	fmt.Println("notes count:", nb.Count())
	fmt.Println(nb.Summary())

	nb.Close()
	nb.Add("this will not be added")
}
