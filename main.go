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

func routers() *chi.Mux {
	
	router.Get("/posts", ViewAllPosts)
	router.Get("/posts/{id}", ViewPost)
	router.Post("/posts", CreatePost)
	router.Put("/posts/{id}", UpdatePost)
	router.Delete("/posts/{id}", DeletePost)

	return router
}

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
	r.Get("/", postHandler.Fetch)
	r.Get("/{id:[0-9]+}", postHandler.GetByID)
	r.Post("/", postHandler.Create)
	r.Put("/{id:[0-9]+}", postHandler.Update)
	r.Delete("/{id:[0-9]+}", postHandler.Delete)

	return r
}

