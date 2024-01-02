package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Parse(filepath string) (Template, error) {
	htmlTpl, err := template.ParseFiles(filepath)
	if err != nil {
		// log.Printf("Parsing template %v", err)
		//	http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl,
	}, nil
}

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html  ,  charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("Problem Executing a Template: %v\n", err)
		log.Printf("Check the views package for errors")
		return
	}
}
