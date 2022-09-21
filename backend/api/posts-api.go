package api

import (
	"net/http"

	"github.com/MrJSdev/go-wp-vue/helpers"
	"github.com/MrJSdev/go-wp-vue/repository"
	"github.com/gorilla/mux"
)

// GetPosts to return all posts from DB
func GetPosts(res http.ResponseWriter, _ *http.Request) {
	posts, err := repository.FindAllPosts()
	if err != nil {
		helpers.ResponseWithError(res, http.StatusBadRequest, err.Error())
		return
	}

	helpers.ResponseWithJSON(res, http.StatusOK, posts)
}

// GetPostByID to return post from DB
func GetPostByID(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	post, err := repository.FindPostByID(params["id"])
	if err != nil {
		helpers.ResponseWithError(res, http.StatusBadRequest, err.Error())
		return
	}

	helpers.ResponseWithJSON(res, http.StatusOK, post)
}
