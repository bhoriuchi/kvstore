package kvstore

import (
	"context"
	"encoding/json"

	"github.com/blevesearch/bleve"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// IndexUpdate an update to an index
type IndexUpdate struct {
	B string // bucket
	K string // Key
	D bool   // Delete
}

// Bucket a bucket
type Bucket struct {
	name  string
	store *Store
	i     bleve.Index
}

// Get gets a value and unmarshals it
func (b *Bucket) Get(ctx context.Context, key string, value interface{}) (err error) {
	var v []byte
	if v, err = b.store.b.Get(ctx, b.name, key); err != nil {
		return
	}

	if err = json.Unmarshal(v, value); err != nil {
		return
	}

	err = b.i.Index(key, value)
	return
}

// Del deletes a value
func (b *Bucket) Del(ctx context.Context, key string, value interface{}) (err error) {
	var data []byte
	if data, err = b.store.b.Del(ctx, b.name, key); err != nil {
		return
	}

	if err = json.Unmarshal(data, value); err != nil {
		return
	}

	err = b.i.Delete(key)
	return
}

// Put marshals the value to json before storing
func (b *Bucket) Put(ctx context.Context, key string, value interface{}) (err error) {
	var data []byte
	if data, err = json.Marshal(value); err != nil {
		return
	}

	if _, err = b.store.b.Put(ctx, b.name, key, data); err != nil {
		return
	}

	err = b.i.Index(string(key), value)
	return
}

// PutProto puts a protobuf message
func (b *Bucket) PutProto(ctx context.Context, key string, value proto.Message) (err error) {
	var data []byte
	jsm := protojson.MarshalOptions{
		AllowPartial:  true,
		UseProtoNames: true,
	}

	if data, err = jsm.Marshal(value); err != nil {
		return
	}

	if _, err = b.store.b.Put(ctx, b.name, key, data); err != nil {
		return
	}

	err = b.i.Index(string(key), value)
	return
}
