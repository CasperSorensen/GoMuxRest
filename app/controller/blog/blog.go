package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/go_rest_postgre/model"
	"github.com/gorilla/mux"
)

var contenttype = "Content-type"
var appJson = "application/json"

func GetAllPosts(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set(contenttype, appJson)

	posts, err := model.GetAllPosts()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(writer).Encode(posts)
	}
}

func GetPost(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set(contenttype, appJson)

	param := mux.Vars(reader)["id"]
	id, err := strconv.ParseUint(param, 10, 64)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	post, err := model.GetPost(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(writer).Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var post model.Article
	err := decoder.Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = model.CreatePost(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var post model.Article
	err := decoder.Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = model.UpdatePost(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)
	idStr := param["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = model.DeletePost(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
