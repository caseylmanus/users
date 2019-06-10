package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := sql.Open("mysql", "app:appsecret@tcp(db:3306)/planet")
	if err != nil {
		logrus.Fatal(err)
	}
	handler := restHandler{db}
	r := chi.NewRouter()
	r.Get("/users", handler.getUsers)
	r.Get("/users/{id}", handler.getUser)
	r.Post("/users", handler.postUser)
	r.Put("/users/{id}", handler.putUser)
	r.Delete("/users/{id}", handler.deleteUser)
	r.Get("/groups", handler.getGroups)
	r.Get("/groups/{id}", handler.getGroup)
	r.Post("/groups", handler.postGroup)
	r.Put("/groups/{id}", handler.putGroup)
	r.Delete("/groups/{id}", handler.deleteGroup)
	logrus.Infoln("Listening on port :8080")
	http.ListenAndServe(":8080", r)

}
