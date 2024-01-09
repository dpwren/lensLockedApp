package main

import (
	"fmt"
	"github.com/dpwren/lensLockedApp/controllers"
	"github.com/dpwren/lensLockedApp/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"path/filepath"
)

type Names struct {
	Name       string
	Newsletter bool
}
type User struct {
	Name   []Names
	Bio    string
	Age    int
	Gender string
}

// func noPageFound(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	_, err := fmt.Fprint(w, "<h1>You have reached a page not found</h1>")
// 	if err != nil {
// 		return
// 	}
// }

// Begin MAIN App
func main() {

	r := chi.NewRouter()
	// middleware stack begins
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))

	// tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	// if err != nil {
	// 	panic(err)
	// }
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err := views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "faqs.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faqs", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "You have reached a page no longer valid", http.StatusInternalServerError)
	})

	// Starting Server beyond this point
	fmt.Println("Server is starting up on Port 3000 ...")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
