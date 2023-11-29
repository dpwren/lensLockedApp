package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	_, err := fmt.Fprint(w, "<h1> Welcome to Davids Photo Buckets </h1>")
	if err != nil {
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := fmt.Fprint(w, "<h1>Contact Page</h1><p> To get in touch, email me at <a href=\"mailto:dpwren@gmail.com\">dpwren@gmail.com</a>.")
	if err != nil {
		return
	}
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	fmt.Fprint(w, `
<style>body{background-color:white;}
	div{background-color:white;align:center}
	ol{background-color:white;}
	.center{
	  margin: auto;
	  width: 60%;
	  border: 3px solid lightblue;
	  padding: 10px;
	} h1{margin: auto;
	  width: 40%;
	  border: 3px solid lightblue;
	  padding: 10px;}
</style>
<body><div class="center"><h1>FAQS Page</h1>
	 <ol>
	<li><strong>Q</strong>  Is there a free version? <br>
     <strong>A</strong> Yes, we offer a free trial for 30 days
    </li>
	<br>
     <li><strong>Q</strong> What are your support hours? <br>
     <strong>A</strong> We have support staff answering calls 24/7
    </li>
	<br>
    <li><strong>Q</strong>How do I contact the support team <br>
     <strong>A</strong> Email support at dpwren@gmail.com
    </li>
	</ol>
</div></body>
	`)
}

func noPageFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := fmt.Fprint(w, "<h1>You have reached a page not found</h1>")
	if err != nil {
		return
	}
}

//func pathHandler(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		homeHandler(w, r)
//	case "/contact":
//		contactHandler(w, r)
//	default:
//		http.Error(w, "Dude! No Page Here", http.StatusNotFound)
//	}
//}

//type Router struct{}
//func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		homeHandler(w, r)
//	case "/contact":
//		contactHandler(w, r)
//	case "/faq":
//		faqHandler(w, r)
//	default:
//		http.Error(w, "Dude! No Page Here", http.StatusNotFound)
//	}
//}

// Begin MAIN App
func main() {
	r := chi.NewRouter()
	//middleware stack begins
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to PhotoBucket"))
	})
	r.Get("/contact", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Contact Page"))
	})

	fmt.Println("Server is starting up on Port 3000 ...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
