package db

import (
	"sync"
	"time"
)

type Engine struct {
	mu     sync.RWMutex
	data   map[string]string
	expire map[string]time.Time
}

func NewEngine() *Engine {
	e := &Engine{
		data:   make(map[string]string),
		expire: make(map[string]time.Time),
	}
	go e.cleanupExpired()
	return e
}

func (e *Engine) Set(key, val string, ttl int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.data[key] = val
	if ttl > 0 {
		e.expire[key] = time.Now().Add(time.Duration(ttl) * time.Second)
	} else {
		delete(e.expire, key)
	}
}

func (e *Engine) Get(key string) (string, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if exp, ok := e.expire[key]; ok && time.Now().After(exp) {
		return "", false
	}
	val, ok := e.data[key]
	return val, ok
}

func (e *Engine) Del(key string) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	_, ok := e.data[key]
	delete(e.data, key)
	delete(e.expire, key)
	return ok
}

func (e *Engine) Expire(key string, seconds int) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, ok := e.data[key]; !ok {
		return false
	}
	e.expire[key] = time.Now().Add(time.Duration(seconds) * time.Second)
	return true
}

func (e *Engine) TTL(key string) int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if exp, ok := e.expire[key]; ok {
		ttl := int(time.Until(exp).Seconds())
		if ttl < 0 {
			return -2
		}
		return ttl
	}
	if _, ok := e.data[key]; ok {
		return -1
	}
	return -2
}

func (e *Engine) cleanupExpired() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		now := time.Now()
		e.mu.Lock()
		for k, exp := range e.expire {
			if now.After(exp) {
				delete(e.data, k)
				delete(e.expire, k)
			}
		}
		e.mu.Unlock()
	}
}
