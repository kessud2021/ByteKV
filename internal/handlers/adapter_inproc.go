package handlers

import (
	"awesomeProject/internal/services"
	"context"
)

// Adapter for in-process usage
type InProcAdapter struct {
	KV *services.KVService
}

func NewInProcAdapter(kv *services.KVService) *InProcAdapter {
	return &InProcAdapter{KV: kv}
}

func (a *InProcAdapter) Set(key, val string) error {
	return a.KV.Set(context.Background(), key, val)
}

func (a *InProcAdapter) Get(key string) (string, error) {
	return a.KV.Get(context.Background(), key)
}

func (a *InProcAdapter) Del(key string) (bool, error) {
	return a.KV.Del(context.Background(), key)
}
