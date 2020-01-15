package main

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-sql-driver/mysql"
)

var router *chi.Mux
var db *sql.DB

const (
	dbHost = "localhost"
	dbPort = "33066"
	dbName = "go-mysql-crud"
	dbPass = "abc123"
)


func main() {
	connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	postHandler := ph.NewPostHandler(connection)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/posts", postRouter(postHandler))
	})

	fmt.Println("Server listen at :8005")
	http.ListenAndServe(":8005", r)
}

func postRouter(postHandler *ph.Post) http.Handler {
	r := chi.NewRouter()
	r.Get("/", postHandler.ViewAllPosts)
	r.Get("/{id:[0-9]+}", postHandler.ViewPost)
	r.Post("/", postHandler.CreatePost)
	r.Put("/{id:[0-9]+}", postHandler.UpdatePost)
	r.Delete("/{id:[0-9]+}", postHandler.DeletePost)

	return r
}

