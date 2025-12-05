package db

import (
	"log"
)

// Example helper to start an in-memory server
func StartServer(addr string) *Listener {
	engine := NewEngine(16)
	listener := NewListener(addr, engine)
	go func() {
		if err := listener.Run(); err != nil {
			log.Println("server stopped:", err)
		}
	}()
	return listener
}
