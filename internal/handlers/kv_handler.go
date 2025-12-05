package handlers

import (
	"encoding/json"
	"net/http"

	"awesomeProject/internal/services"

	"github.com/go-chi/chi/v5"
)

// RegisterKVRoutes registers KV HTTP routes for the chi router
func RegisterKVRoutes(r chi.Router, kvSvc *services.KVService) {
	r.Post("/kv", setHandler(kvSvc))
	r.Get("/kv/{key}", getHandler(kvSvc))
	r.Delete("/kv/{key}", delHandler(kvSvc))
}

// setHandler handles POST /kv
func setHandler(kv *services.KVService) http.HandlerFunc {
	type req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var body req
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if err := kv.Set(r.Context(), body.Key, body.Value); err != nil {
			http.Error(w, "failed to set value", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ok":true}`))
	}
}

// getHandler handles GET /kv/{key}
func getHandler(kv *services.KVService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")
		val, err := kv.Get(r.Context(), key)
		if err != nil {
			http.Error(w, "key not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"key":   key,
			"value": val,
		})
	}
}

// delHandler handles DELETE /kv/{key}
func delHandler(kv *services.KVService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")
		_, _ = kv.Del(r.Context(), key)
		w.WriteHeader(http.StatusOK)
	}
}
