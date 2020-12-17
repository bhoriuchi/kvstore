package kvstore

import (
	"github.com/bhoriuchi/kvstore/backend"
	"github.com/blevesearch/bleve"
)

// Store a store
type Store struct {
	b       backend.Backend
	buckets map[string]*Bucket
}

// Bucket creates or returns a bucket
func (s *Store) Bucket(name string) (bucket *Bucket, err error) {
	if s.buckets == nil {
		s.buckets = map[string]*Bucket{}
	}

	var ok bool
	if bucket, ok = s.buckets[name]; ok {
		return
	}

	bucket = &Bucket{
		name:  name,
		store: s,
	}

	bucket.i, err = bleve.Open(name)
	if err == nil {
		return
	}

	if err != bleve.ErrorIndexPathDoesNotExist {
		return
	}

	err = nil
	mapping := bleve.NewIndexMapping()
	if bucket.i, err = bleve.New(name, mapping); err != nil {
		return
	}

	s.buckets[name] = bucket
	return
}

// NewStore creates a new kvstore instance
func NewStore(b backend.Backend) (store *Store) {
	store = &Store{
		b:       b,
		buckets: map[string]*Bucket{},
	}
	return
}
