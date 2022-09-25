package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MrJSdev/go-wp-vue/entity"
	"github.com/MrJSdev/go-wp-vue/helpers"
	"github.com/MrJSdev/go-wp-vue/repository"
	"github.com/gorilla/mux"
)

// GetCommentsByPostID to return comments list from DB
func GetCommentsByPostID(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["post-id"])
	comments, err := repository.FindCommentsByPostID(id)

	if err != nil {
		helpers.ResponseWithError(res, http.StatusBadRequest, err.Error())
		return
	}

	helpers.ResponseWithJSON(res, http.StatusOK, comments)
}

func SaveCommentByPostID(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["post-id"])
	var comment = entity.Comment{}

	err := json.NewDecoder(req.Body).Decode(&comment)

	if err != nil {
		helpers.ResponseWithError(res, http.StatusBadRequest, err.Error())
		return
	}

	err2 := repository.AddNewCommentByPostID(id, comment)

	if err2 != nil {
		helpers.ResponseWithError(res, http.StatusBadRequest, err2.Error())
		return
	}

	helpers.ResponseWithJSON(res, http.StatusOK, comment)
}
