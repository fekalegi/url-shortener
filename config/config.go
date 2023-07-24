package config

import (
	"os"
	"strconv"
)

type InitManager interface {
	InitConfigMemCache()
	GetMemCacheConfig() *memCacheConfig
}

type memCacheConfig struct {
	Host string
	Port int
}

func (m *memCacheConfig) GetMemCacheConfig() *memCacheConfig {
	return m
}

func NewMemCacheConfig() InitManager {
	newConfig := new(memCacheConfig)
	newConfig.InitConfigMemCache()
	return newConfig
}

func (m *memCacheConfig) InitConfigMemCache() {
	port, _ := strconv.Atoi(os.Getenv("MEMCACHE_PORT"))
	m.Host = os.Getenv("MEMCACHE_HOST")
	m.Port = port
}
