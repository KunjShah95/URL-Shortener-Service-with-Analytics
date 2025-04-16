package storage

// Post is defined in memory.go

type PostStore interface {
	GetAll() ([]Post, error)
	Get(id int) (*Post, error)
	Create(post *Post) error
	Delete(id int) error
}