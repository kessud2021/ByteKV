// internal/store/store.go
package store

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type TCPStore struct {
	addr string
}

func NewTCPStore(addr string) *TCPStore {
	return &TCPStore{addr: addr}
}

// --- internal roundtrip to TCP server ---
func (s *TCPStore) roundtrip(ctx context.Context, req string) (string, error) {
	d := &net.Dialer{}
	conn, err := d.DialContext(ctx, "tcp", s.addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(2 * time.Second))
	_, err = conn.Write([]byte(req))
	if err != nil {
		return "", err
	}

	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	r := bufio.NewReader(conn)
	return readReply(r)
}

// --- RESP protocol builder ---
func BuildRESP(args ...string) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("*%d\r\n", len(args)))
	for _, a := range args {
		b.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(a), a))
	}
	return b.String()
}

// --- read RESP reply ---
func readReply(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return "", nil
	}

	switch line[0] {
	case '+':
		return strings.TrimPrefix(line, "+"), nil
	case '-':
		return "", fmt.Errorf("ERR %s", strings.TrimPrefix(line, "-"))
	case ':':
		return strings.TrimPrefix(line, ":"), nil
	case '$':
		n, _ := strconv.Atoi(strings.TrimPrefix(line, "$"))
		if n < 0 {
			return "", nil
		}
		buf := make([]byte, n+2)
		_, err := r.Read(buf)
		if err != nil {
			return "", err
		}
		return string(buf[:n]), nil
	case '*':
		n, _ := strconv.Atoi(strings.TrimPrefix(line, "*"))
		parts := make([]string, 0, n)
		for i := 0; i < n; i++ {
			item, err := readReply(r)
			if err != nil {
				return "", err
			}
			parts = append(parts, item)
		}
		return strings.Join(parts, " "), nil
	default:
		return line, nil
	}
}

// --- Convenience methods for TCPStore ---
func (s *TCPStore) Set(ctx context.Context, key, val string) error {
	req := BuildRESP("SET", key, val)
	_, err := s.roundtrip(ctx, req)
	return err
}

func (s *TCPStore) Get(ctx context.Context, key string) (string, error) {
	req := BuildRESP("GET", key)
	resp, err := s.roundtrip(ctx, req)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func (s *TCPStore) Del(ctx context.Context, key string) (bool, error) {
	req := BuildRESP("DEL", key)
	resp, err := s.roundtrip(ctx, req)
	if err != nil {
		return false, err
	}
	n, _ := strconv.Atoi(resp)
	return n > 0, nil
}

func (s *TCPStore) Publish(ctx context.Context, channel, msg string) (int, error) {
	req := BuildRESP("PUBLISH", channel, msg)
	resp, err := s.roundtrip(ctx, req)
	if err != nil {
		return 0, err
	}
	n, _ := strconv.Atoi(resp)
	return n, nil
}
