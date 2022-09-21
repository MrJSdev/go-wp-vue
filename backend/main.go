package main

import (
	"net/http"
	"time"

	"github.com/MrJSdev/go-wp-vue/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./frontend/.output/public"))
	http.Handle("/", fs)
	// router.HandleFunc("/", fs.ServeHTTP)
	router.HandleFunc("/api/posts", api.GetPosts)
	router.HandleFunc("/api/posts/{id}", api.GetPostByID)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8081",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
