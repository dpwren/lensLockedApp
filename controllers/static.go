package controllers

import (
	"github.com/dpwren/lensLockedApp/views"
	"html/template"
	"net/http"
)

type Static struct {
	Template views.Template
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes, at this time we offer a 30 day trial.",
		},
		{
			Question: "What are your support hours?",
			Answer:   " We have support staff answering calls 24/7",
		},
		{
			Question: "How do I contact the support team",
			Answer:   `Email Us - <a href="mailto:dpwren@gmail.com">dpwren@gmail.com</a>`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
