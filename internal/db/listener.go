package db

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

type Listener struct {
	addr   string
	engine *Engine
	ln     net.Listener
	quit   chan struct{}
}

func NewListener(addr string, engine *Engine) *Listener {
	return &Listener{addr: addr, engine: engine, quit: make(chan struct{})}
}

func (l *Listener) Run() error {
	ln, err := net.Listen("tcp", l.addr)
	if err != nil {
		return err
	}
	l.ln = ln

	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-l.quit:
				return nil
			default:
				continue
			}
		}
		go l.handleConn(conn)
	}
}

func (l *Listener) Stop(ctx context.Context) {
	close(l.quit)
	if l.ln != nil {
		_ = l.ln.Close()
	}
}

func (l *Listener) handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)

	for {
		conn.SetReadDeadline(time.Now().Add(300 * time.Second))
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			conn.Write([]byte("-ERR " + err.Error() + "\r\n"))
			return
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		args := strings.Split(line, " ")
		cmd := strings.ToUpper(args[0])

		switch cmd {
		case "PING":
			conn.Write([]byte("+PONG\r\n"))
		case "SET":
			if len(args) < 3 {
				conn.Write([]byte("-ERR wrong args\r\n"))
				continue
			}
			key := args[1]
			val := args[2]
			ttl := 0
			if len(args) >= 4 {
				ttl, _ = strconv.Atoi(args[3])
			}
			l.engine.Set(key, val, ttl)
			conn.Write([]byte("+OK\r\n"))
		case "GET":
			if len(args) < 2 {
				conn.Write([]byte("-ERR wrong args\r\n"))
				continue
			}
			key := args[1]
			v, ok := l.engine.Get(key)
			if !ok {
				conn.Write([]byte("$-1\r\n"))
			} else {
				conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(v), v)))
			}
		case "DEL":
			if len(args) < 2 {
				conn.Write([]byte("-ERR wrong args\r\n"))
				continue
			}
			ok := l.engine.Del(args[1])
			conn.Write([]byte(fmt.Sprintf(":%d\r\n", boolToInt(ok))))
		case "EXPIRE":
			if len(args) < 3 {
				conn.Write([]byte("-ERR wrong args\r\n"))
				continue
			}
			sec, _ := strconv.Atoi(args[2])
			ok := l.engine.Expire(args[1], sec)
			conn.Write([]byte(fmt.Sprintf(":%d\r\n", boolToInt(ok))))
		case "TTL":
			if len(args) < 2 {
				conn.Write([]byte("-ERR wrong args\r\n"))
				continue
			}
			ttl := l.engine.TTL(args[1])
			conn.Write([]byte(fmt.Sprintf(":%d\r\n", ttl)))
		case "QUIT":
			conn.Write([]byte("+OK\r\n"))
			return
		default:
			conn.Write([]byte("-ERR unknown command\r\n"))
		}
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
