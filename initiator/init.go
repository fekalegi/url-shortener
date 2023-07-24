package initiator

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-gonic/gin"
	"url-shortener/config"
)

type InitiationManager interface {
	initGin()
	initMemCache()

	GetMemCache() *memcache.Client
	GetGin() *gin.Engine
}

type initiator struct {
	gin            *gin.Engine
	memCacheClient *memcache.Client
}

func (i *initiator) GetGin() *gin.Engine {
	return i.gin
}

func (i *initiator) GetMemCache() *memcache.Client {
	return i.memCacheClient
}

func NewInit() InitiationManager {
	initiation := new(initiator)
	initiation.initMemCache()
	initiation.initGin()
	return initiation
}

func (i *initiator) initGin() {
	i.gin = gin.Default()
}

func (i *initiator) initMemCache() {
	conf := config.NewMemCacheConfig().GetMemCacheConfig()
	host := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	i.memCacheClient = memcache.New(host)
}
