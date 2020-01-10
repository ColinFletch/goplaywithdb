package main

import (
	"database/sql"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

var router *chi.Mux
var db *sql.DB

const (
	dbHost = "localhost"
	dbPort = "33066"
	dbName = "go-mysql-crud"
	dbPass = "abc123"
)

func routers() *chi.Mux {
	router.Get("/posts", AllPosts)
	router.Get("/posts/{id}", ViewPost)
	router.Post("/posts", CreatePost)
	router.Put("/posts/{id}", UpdatePost)
	router.Delete("/posts/{id}", DeletePost)

	return router
}
