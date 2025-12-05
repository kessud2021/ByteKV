package services

import (
	"awesomeProject/internal/store"
	"context"
	"log"
)

type KVService struct {
	store  *store.TCPStore
	logger *log.Logger
}

func NewKVService(s *store.TCPStore, l *log.Logger) *KVService {
	return &KVService{store: s, logger: l}
}

func (k *KVService) Set(ctx context.Context, key, val string) error {
	return k.store.Set(ctx, key, val)
}

func (k *KVService) Get(ctx context.Context, key string) (string, error) {
	return k.store.Get(ctx, key)
}

func (k *KVService) Del(ctx context.Context, key string) (bool, error) {
	return k.store.Del(ctx, key)
}
