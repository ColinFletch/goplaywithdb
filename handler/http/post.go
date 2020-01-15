package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	models "github.com/colinfletch/goplaywithdb/model"
	repository "github.com/colinfletch/goplaywithdb/repository"
	post "github.com/colinfletch/goplaywithdb/repository/postrepo"

	"github.com/colinfletch/goplaywithdb/driver"
	"github.com/go-chi/chi"
)

// NewPostHandler ...
func NewPostHandler(db *driver.DB) *Post {
	return &Post{
		repo: post.NewSQLPostRepo(db.SQL),
	}
}

// Post ...
type Post struct {
	repo repository.PostRepo
}

// ViewAllPosts all post data
func (p *Post) ViewAllPosts(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.ViewAllPosts(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}

// CreatePost a new post
func (p *Post) CreatePost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	json.NewDecoder(r.Body).Decode(&post)
	newID, err := p.repo.CreatePost(r.Context(), &post)
	fmt.Println(newID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// UpdatePost a post by id
func (p *Post) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := models.Post{ID: int32(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := p.repo.UpdatePost(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// ViewPost returns a post details
func (p *Post) ViewPost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := p.repo.ViewPost(r.Context(), int32(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete a post
func (p *Post) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := p.repo.DeletePost(r.Context(), int32(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
