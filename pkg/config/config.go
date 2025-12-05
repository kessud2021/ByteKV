package config

import "os"

type Config struct {
	LogLevel   string
	DBShards   int
	AOFPath    string
	BindDBAddr string
}

func Load() Config {
	return Config{
		LogLevel:   getenv("LOG_LEVEL", "debug"),
		DBShards:   1,
		AOFPath:    getenv("AOF_PATH", "data.aof"),
		BindDBAddr: getenv("DB_ADDR", ":6380"),
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
