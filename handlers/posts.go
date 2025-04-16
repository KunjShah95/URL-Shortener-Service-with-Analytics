package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/KunjShah95/URLShortenerServicewithAnalytics/storage"
)

type PostHandler struct {
	Store storage.PostStore
}

func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Store.GetAll()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Error reading body")
		return
	}

	var post storage.Post
	if err := json.Unmarshal(body, &post); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if post.Body == "" {
		sendError(w, http.StatusBadRequest, "Body cannot be empty")
		return
	}

	if err := h.Store.Create(&post); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) GetOne(w http.ResponseWriter, r *http.Request, id int) {
	post, err := h.Store.Get(id)
	if err != nil {
		if err == storage.ErrNotFound {
			sendError(w, http.StatusNotFound, "Post not found")
		} else {
			sendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.Store.Delete(id); err != nil {
		if err == storage.ErrNotFound {
			sendError(w, http.StatusNotFound, "Post not found")
		} else {
			sendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func sendError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"error":   http.StatusText(code),
		"message": message,
	})
}