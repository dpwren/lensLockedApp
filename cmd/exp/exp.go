package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type User struct {
	Name       string
	Bio        string
	Age        int
	Gender     string
	Newsletter bool
}

func executeTemplate(w http.ResponseWriter, tplPath string) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("Parsing template %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		fmt.Printf("Nil # code is: %v", nil)
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func main() {
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
