package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/skurtz97/splat/internal/splat"
)

func (s *Server) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(`{"message": "alive"}`))
}

func (s *Server) ListPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	posts, err := s.service.ListPosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error getting posts", "error:" : "` + err.Error() + `"}`))
		return
	}

	postsJson, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error marshalling posts", "error:" : "` + err.Error() + `"}`))
		return
	}

	w.Write(postsJson)
}

func (s *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := chi.URLParam(r, "id")
	post, err := s.service.GetPost(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error getting post", "error:" : "` + err.Error() + `"}`))
		return
	}

	postJson, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error marshalling post", "error:" : "` + err.Error() + `"}`))
		return
	}

	w.Write(postJson)
}

func (s *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	postBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error reading post", "error:" : "` + err.Error() + `"}`))
		return
	}

	var post splat.Post
	err = json.Unmarshal(postBytes, &post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error marshalling post", "error:" : "` + err.Error() + `"}`))
		return
	}

	if err := s.service.CreatePost(&post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error creating post", "error:" : "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(postBytes)
}

func (s *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	postBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error reading post", "error:" : "` + err.Error() + `"}`))
		return
	}

	var post splat.Post
	err = json.Unmarshal(postBytes, &post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error marshalling post", "error:" : "` + err.Error() + `"}`))
		return
	}

	if err := s.service.UpdatePost(&post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error updating post", "error:" : "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(postBytes)
}

func (s *Server) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := chi.URLParam(r, "id")

	if err := s.service.DeletePost(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error deleting post", "error:" : "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
}
