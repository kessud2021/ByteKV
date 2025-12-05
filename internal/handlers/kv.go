package handlers

import "awesomeProject/internal/services"

type KVHandler struct {
	KV *services.KVService
}

func NewKVHandler(kv *services.KVService) *KVHandler {
	return &KVHandler{KV: kv}
}
