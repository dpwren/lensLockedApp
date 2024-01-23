package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

/*
MUST function for func Parse (TEMPLATE, ERROR)
*/

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParsePS(fs fs.FS, patterns ...string) (Template, error) {
	htmlTpl, err := template.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl,
	}, nil
}

/*
 func Parse returns a Template(){htmlTpl,nil}
*/

func Parse(filepath string) (Template, error) {
	htmlTpl, err := template.ParseFiles(filepath)
	if err != nil {
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
	// w.Header().Set("Content-Type", "text/html  ,  charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("Problem executing a Template: %v\n", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}
