package backend

import (
	"context"
	"errors"
)

// Errors
var (
	ErrBucketNotFound = errors.New("ERR_BUCKET_NOT_FOUND")
	ErrEntryNotFound  = errors.New("ERR_ENTRY_NOT_FOUND")
)

// Meta metadata
type Meta struct {
	Created   int64
	Updated   int64
	Protected int32
}

// Entry an entry
type Entry struct {
	Key  string
	Meta *Meta
	Data []byte
}

// Backend the backend interface
type Backend interface {
	Get(ctx context.Context, bucket, key string) (out []byte, err error)
	Del(ctx context.Context, bucket, key string) (out []byte, err error)
	Put(ctx context.Context, bucket, key string, in []byte) (out []byte, err error)
	List(ctx context.Context, bucket string) (list [][]byte, err error)
}
