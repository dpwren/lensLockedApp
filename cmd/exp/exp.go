package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"os"
)

type User struct {
	Name       string
	Bio        string
	Age        int
	Gender     string
	Newsletter bool
}

func main() {

	if _, err := os.Open(""); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		}
		if errors.Is(err, fs.ErrInvalid) {
			fmt.Printf(">>>> %q is not valid", err)
		}
	}

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name:       "John Smith",
		Bio:        "I am John Smith. I like ice cream and play soccer.",
		Age:        44,
		Gender:     "Male",
		Newsletter: false,
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
