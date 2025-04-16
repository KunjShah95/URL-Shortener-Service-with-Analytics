package main

import (
	"go-server/handlers"
	"go-server/middleware"
	"go-server/storage"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// Initialize storage
	store := storage.NewMemoryStore()
	postHandler := &handlers.PostHandler{Store: store}

	// Setup routes with middleware chain
	http.Handle("/posts", middleware.Logging(middleware.JSON(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			postHandler.GetAll(w, r)
		case http.MethodPost:
			postHandler.Create(w, r)
		default:
			handlers.SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	}))))

	http.Handle("/posts/", middleware.Logging(middleware.JSON(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Path[len("/posts/"):])
		if err != nil {
			handlers.SendError(w, http.StatusBadRequest, "Invalid post ID")
			return
		}

		switch r.Method {
		case http.MethodGet:
			postHandler.GetOne(w, r, id)
		case http.MethodDelete:
			postHandler.Delete(w, r, id)
		default:
			handlers.SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	}))))

	http.Handle("/health", middleware.Logging(http.HandlerFunc(handlers.HealthHandler)))

	// Server configuration
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Start server
	log.Printf("Server is running at http://localhost%s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
