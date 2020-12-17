package kvstore

import (
	"context"
	"testing"

	"github.com/bhoriuchi/kvstore/backend/memory"
)

func TestCrud(t *testing.T) {
	ctx := context.Background()
	backend := memory.NewBackend()
	kv := NewStore(backend)

	testValues := []string{
		"foo",
		"bar",
	}

	buc, err := kv.Bucket("foo")
	if err != nil {
		t.Error(err)
		return
	}

	for _, value := range testValues {
		if err = buc.Put(ctx, "/test", value); err != nil {
			t.Error(err)
			return
		}

		var (
			storedValue  interface{}
			deletedValue interface{}
		)
		buc.Get(ctx, "/test", &storedValue)
		t.Logf("Comparing %v = %v", value, storedValue)
		if value != storedValue {
			t.Error("get value does not match put value")
			return
		}

		if err = buc.Del(ctx, "/test", &deletedValue); err != nil {
			t.Error(err)
			return
		}

		if deletedValue != storedValue {
			t.Error("deleted value does not match stored value")
			return
		}
	}
}
