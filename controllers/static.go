package controllers

import (
	"github.com/dpwren/lensLockedApp/views"
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
