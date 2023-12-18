package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	tplPath := filepath.Join("templates", "home.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("Parsing template %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", "contact.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("Parsing template %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	tplPath := filepath.Join("templates", "faqs.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("Parsing template %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}

}

func noPageFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := fmt.Fprint(w, "<h1>You have reached a page not found</h1>")
	if err != nil {
		return
	}
}

// Begin MAIN App
func main() {
	r := chi.NewRouter()
	// middleware stack begins
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faqs", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
	})

	fmt.Println("Server is starting up on Port 3000 ...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
