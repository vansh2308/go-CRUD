package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vansh2308/go-CHI-CRUD.git/controllers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {
		r.Get("/", controllers.GetArticle)
		r.Post("/", controllers.CreateArticle)
		r.Put("/{id}", controllers.UpdateArticle)
		r.Put("/{id}", controllers.UpdateArticle)
		r.Delete("/{id}", controllers.DeleteArticle)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})

	log.Println("Server running at http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
