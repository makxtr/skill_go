package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Необходимо реализовать in-memory cache, который реализует операции добавления
// поиска и удаления элемента за минимальное время.
// Кэш должен быть конкуретнтно-безопасным, пригодным для использования в многопоточной среде

// добавить TTL(каждый элемент должен иметь время жизни)

var ErrNotFound = errors.New("key not found")

type ICache interface {
	Get(context.Context, string) (string, error)
	Set(context.Context, string) (string, error)
	Del(context.Context, string) error
}

type elem struct {
	val     string
	expDate time.Time
}

type Cache struct {
	storage map[string]elem
	mu      *sync.RWMutex
	ttl     time.Duration
	done    chan struct{}
}

func New(ttl time.Duration) *Cache {
	cache := &Cache{
		storage: make(map[string]elem),
		mu:      &sync.RWMutex{},
		ttl:     ttl,
		done:    make(chan struct{}),
	}

	cache.clearByTTL()

	return cache
}

func (c *Cache) clearByTTL() {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				c.clear()
			case <-c.done:
				return
			}
		}
	}()
}

func (c *Cache) Stop() {
	close(c.done)
}

func (c *Cache) clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, el := range c.storage {
		if el.expDate.Before(time.Now()) {
			delete(c.storage, key)
		}
	}
}

func (c *Cache) Get(_ context.Context, key string) (string, error) {
	c.mu.RLock()
	el, ok := c.storage[key]
	c.mu.RUnlock()

	if !ok {
		return "", ErrNotFound
	}
	if el.expDate.Before(time.Now()) {
		c.del(key)
		return "", ErrNotFound
	}

	return el.val, nil
}

func (c *Cache) Set(_ context.Context, key string, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	el := elem{
		val:     value,
		expDate: time.Now().Add(c.ttl),
	}

	c.storage[key] = el

	return nil
}

func (c *Cache) del(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.storage, key)
}

func (c *Cache) Del(_ context.Context, key string) error {
	c.del(key)

	return nil
}
