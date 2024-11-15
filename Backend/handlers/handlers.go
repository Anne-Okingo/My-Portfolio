package handlers

import (
	"net/http"

	"my-portfolio/Backend/renders"
)

// HomeHandler handles the homepage route '/'
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renders.RenderTemplate(w, "index.html", nil)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// NotFoundHandler handles unknown routes; 404 status
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renders.RenderTemplate(w, "notfound.page.html", nil)
}

// BadRequestHandler handles bad requests routes
func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	renders.RenderTemplate(w, "badrequest.page.html", nil)
}

// ServerErrorHandler handles server failures that result in status 500
func ServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	renders.RenderTemplate(w, "servererror.page.html", nil)
}

// package handlers

// import (
// 	"my-portfolio/Backend/renders"
// 	"net/http"
// )

// func HomeHandler(w http.ResponseWriter, r*http.Request){
// 	if r.Method == http.MethodGet{
// 		renders.RenderTemplate(w,"index.html",nil)
// 	}else{
// 		http.Error(w, "Method no allowed",http.StatusMethodNotAllowed)
// 	}
// }
