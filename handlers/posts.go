package handlers

import (
	"encoding/json"
	"go-server/storage"
	"io"
	"net/http"
)

type PostHandler struct {
	Store storage.PostStore
}

func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Store.GetAll()
	if err != nil {
		SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		SendError(w, http.StatusBadRequest, "Error reading body")
		return
	}

	var post storage.Post
	if err := json.Unmarshal(body, &post); err != nil {
		SendError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if post.Content == "" {
		SendError(w, http.StatusBadRequest, "Body cannot be empty")
		return
	}

	if err := h.Store.Create(&post); err != nil {
		SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) GetOne(w http.ResponseWriter, r *http.Request, id int) {
	post, err := h.Store.Get(id)
	if err != nil {
		if err == storage.ErrNotFound {
			SendError(w, http.StatusNotFound, "Post not found")
		} else {
			SendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.Store.Delete(id); err != nil {
		if err == storage.ErrNotFound {
			SendError(w, http.StatusNotFound, "Post not found")
		} else {
			SendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// SendError sends a JSON error response with the given status code and message
func SendError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"error":   http.StatusText(code),
		"message": message,
	})
}
