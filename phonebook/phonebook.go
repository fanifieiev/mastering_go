package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		exe := path.Base(args[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	data = append(data, Entry{"Fevzi", "Anif", "1312312313"})
	data = append(data, Entry{"Melek", "Anif", "3545345345345"})
	data = append(data, Entry{"Amira", "Anif", "5353453534"})
	data = append(data, Entry{"Milya", "Anif", "13123121312"})

	switch args[1] {
	case "search":
		if len(args) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(args[2])
		if result == nil {
			fmt.Println("Entry not found", args[2])
			return
		}
		fmt.Println("Result: ", *result)
	case "list":
		list()

	default:
		fmt.Println("Not a valid option")
	}

}
