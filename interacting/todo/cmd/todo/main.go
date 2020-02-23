package main

import (
	"fmt"
	"os"
	"pragprog.com/rggo/interacting/todo"
	"strings"
)

const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	// use the get method to read to do items from the file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
	}

	// save the new list
	if err := l.Save(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

