package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
	"go.uber.org/zap"
	_ "github.com/mattn/go-sqlite3"
)

// Post represents a blog post
type Post struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var (
	posts   = make(map[int]Post)
	nextID  = 1
	postsMu sync.Mutex
)

func main() {
	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Initialize database
	if err := initDB(); err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}
	defer db.Close()

	// Setup routes with middleware chain
	http.Handle("/posts", loggingMiddleware(jsonMiddleware(http.HandlerFunc(postsHandler))))
	http.Handle("/posts/", loggingMiddleware(jsonMiddleware(http.HandlerFunc(postHandler))))

	// Health check endpoint
	http.Handle("/health", loggingMiddleware(http.HandlerFunc(healthHandler)))

	// Server configuration
	server := &http.Server{
		Addr:         getPort(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("Server is running at http://localhost%s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

// Middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Handlers
func postsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetPosts(w, r)
	case http.MethodPost:
		handlePostPosts(w, r)
	default:
		sendError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/posts/"):])
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	switch r.Method {
	case http.MethodGet:
		handleGetPost(w, r, id)
	case http.MethodDelete:
		handleDeletePost(w, r, id)
	default:
		sendError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

// Handler functions
func handleGetPosts(w http.ResponseWriter, r *http.Request) {
	postsMu.Lock()
	defer postsMu.Unlock()

	// Convert map to slice
	postList := make([]Post, 0, len(posts))
	for _, p := range posts {
		postList = append(postList, p)
	}

	json.NewEncoder(w).Encode(postList)
}

func handlePostPosts(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Error reading request body")
		return
	}

	var p Post
	if err := json.Unmarshal(body, &p); err != nil {
		sendError(w, http.StatusBadRequest, "Error unmarshalling JSON")
		return
	}

	// Validate post
	if p.Body == "" {
		sendError(w, http.StatusBadRequest, "Post body cannot be empty")
		return
	}

	postsMu.Lock()
	defer postsMu.Unlock()

	// Set post metadata
	p.ID = nextID
	p.CreatedAt = time.Now()
	p.UpdatedAt = p.CreatedAt

	posts[p.ID] = p
	nextID++

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func handleGetPost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	p, ok := posts[id]
	if !ok {
		sendError(w, http.StatusNotFound, "Post not found")
		return
	}

	json.NewEncoder(w).Encode(p)
}

func handleDeletePost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	_, ok := posts[id]
	if !ok {
		sendError(w, http.StatusNotFound, "Post not found")
		return
	}

	delete(posts, id)
	w.WriteHeader(http.StatusNoContent)
}

// Helper functions
func sendError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"error":   http.StatusText(code),
		"message": message,
	})
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

var db *sql.DB

func initDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./posts.db")
	if err != nil {
		return err
	}
	
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	body TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);`)
	return err
}
