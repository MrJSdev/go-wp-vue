package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MrJSdev/go-wp-vue/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.HandleFunc("/api/posts", api.GetPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", api.GetPostByID).Methods("GET")
	router.HandleFunc("/api/comments/{post-id}", api.GetCommentsByPostID).Methods("GET")
	router.HandleFunc("/api/comments/{post-id}", api.SaveCommentByPostID).Methods("POST")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
