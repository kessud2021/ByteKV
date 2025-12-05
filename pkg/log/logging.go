package logging

import (
	"log"
)

type Logger struct{}

func New(level string) *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string, args ...any) {
	log.Println("[INFO]", msg, args)
}

func (l *Logger) Warn(msg string, args ...any) {
	log.Println("[WARN]", msg, args)
}

func (l *Logger) Fatal(msg string, args ...any) {
	log.Fatal("[FATAL]", msg, args)
}
