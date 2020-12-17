package memory

import (
	"context"

	"github.com/bhoriuchi/kvstore/backend"
)

// Bucket a bucket
type Bucket struct {
	data map[string][]byte
}

// Backend memory backend
type Backend struct {
	c map[string]*Bucket
}

// NewBackend creates a new backend
func NewBackend() (backend *Backend) {
	backend = &Backend{
		c: map[string]*Bucket{},
	}
	return
}

func (b *Backend) bucket(name string) (bucket *Bucket, err error) {
	var found bool
	if bucket, found = b.c[name]; !found {
		bucket = &Bucket{
			data: map[string][]byte{},
		}
		b.c[name] = bucket
		return
	}

	if bucket.data == nil {
		bucket.data = map[string][]byte{}
	}

	return
}

// Get get data
func (b *Backend) Get(ctx context.Context, bucket, key string) (out []byte, err error) {
	var (
		buc   *Bucket
		found bool
	)

	if buc, err = b.bucket(bucket); err != nil {
		return
	}

	if out, found = buc.data[key]; !found {
		err = backend.ErrEntryNotFound
		return
	}

	return
}

// Del deletes the data
func (b *Backend) Del(ctx context.Context, bucket, key string) (out []byte, err error) {
	var (
		buc *Bucket
	)

	if buc, err = b.bucket(bucket); err != nil {
		return
	}

	out = buc.data[key]
	delete(buc.data, key)
	return
}

// Put puts a value
func (b *Backend) Put(ctx context.Context, bucket, key string, in []byte) (out []byte, err error) {
	var buc *Bucket

	if buc, err = b.bucket(bucket); err != nil {
		return
	}

	buc.data[key] = in
	out = in
	return
}

// List lists all data
func (b *Backend) List(ctx context.Context, bucket string) (list [][]byte, err error) {
	var buc *Bucket

	if buc, err = b.bucket(bucket); err != nil {
		return
	}

	list = [][]byte{}
	for _, entry := range buc.data {
		list = append(list, entry)
	}
	return
}
