package storage

import (
	"errors"
	"sync"
	"time"
)

// Post represents a post entity in the storage
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MemoryStore struct {
	posts   map[int]Post
	nextID  int
	postsMu sync.Mutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		posts:  make(map[int]Post),
		nextID: 1,
	}
}

func (s *MemoryStore) GetAll() ([]Post, error) {
	s.postsMu.Lock()
	defer s.postsMu.Unlock()

	postList := make([]Post, 0, len(s.posts))
	for _, p := range s.posts {
		postList = append(postList, p)
	}
	return postList, nil
}

func (s *MemoryStore) Get(id int) (*Post, error) {
	s.postsMu.Lock()
	defer s.postsMu.Unlock()

	post, exists := s.posts[id]
	if !exists {
		return nil, ErrNotFound
	}
	return &post, nil
}

func (s *MemoryStore) Create(post *Post) error {
	s.postsMu.Lock()
	defer s.postsMu.Unlock()

	post.ID = s.nextID
	post.CreatedAt = time.Now()
	post.UpdatedAt = post.CreatedAt
	s.posts[post.ID] = *post
	s.nextID++
	return nil
}

func (s *MemoryStore) Delete(id int) error {
	s.postsMu.Lock()
	defer s.postsMu.Unlock()

	if _, exists := s.posts[id]; !exists {
		return ErrNotFound
	}
	delete(s.posts, id)
	return nil
}

var ErrNotFound = errors.New("not found")