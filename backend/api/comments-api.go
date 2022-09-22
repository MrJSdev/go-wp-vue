package api

import (
	"net/http"

	"github.com/MrJSdev/go-wp-vue/helpers"
	"github.com/MrJSdev/go-wp-vue/repository"
	"github.com/gorilla/mux"
)

// GetCommentsByPostID to return comments list from DB
func GetCommentsByPostID(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	comments, err := repository.FindCommentsByPostID(params["post-id"])

	if err != nil {
		helpers.ResponseWithError(res, http.StatusBadRequest, err.Error())
		return
	}

	helpers.ResponseWithJSON(res, http.StatusOK, comments)
}
