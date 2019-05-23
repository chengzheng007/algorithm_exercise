package main

import (
	"container/list"
	"errors"
)

var (
	// ErrCacheSize indicate cache size must great than zero
	ErrCacheSize = errors.New("invalid lru cache size")
)

type lruCache struct {
	size    int
	hashMap map[string]int
	dblist  *list.List
}

func NewLruCache(size int) (*lruCache, error) {
	if size <= 0 {
		return nil, ErrCacheSize
	}
	return &lruCache{
		size:    size,
		hashMap: make(map[string]int),
		dblist:  list.New(),
	}, nil
}

func (lch *lruCache) Put(k string, v int) {

}
