package main

import (
	"fmt"
	"github.com/dpwren/lensLockedApp/controllers"
	"github.com/dpwren/lensLockedApp/templates"
	"github.com/dpwren/lensLockedApp/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Names struct {
	Name       string
	Newsletter bool
}
type User struct {
	Name   string
	Bio    string
	Age    int
	Gender string
}

// Begin MAIN App
func main() {

	r := chi.NewRouter()
	// middleware stack begins
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml",
		"tailwind.gohtml"))

	// r.Get("/signup", usersC.myNew)

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faqs", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faqs.gohtml", "tailwind.gohtml"))))

	r.Get("/ex", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "signup.gohtml"))))

	r.Get("/signup", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "You have reached a page no longer valid", http.StatusInternalServerError)
	})

	// Starting Server beyond this point
	fmt.Println("Server is starting up on Port 3000 ...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
