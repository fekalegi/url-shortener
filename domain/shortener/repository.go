package shortener

import (
	"encoding/json"
	"errors"
	"github.com/bradfitz/gomemcache/memcache"
	"log"
)

type repository struct {
	db *memcache.Client
}

func NewShortenerRepository(db *memcache.Client) Repository {
	return &repository{
		db,
	}
}

type keysCache struct {
	Keys []string `json:"keys"`
}

var defaultKey = "key"

//go:generate mockgen -destination=../../mocks/repository/mock_shortener_repository.go -package=mock_repository -source=repository.go
type Repository interface {
	Store(req *Link) error
	GetByShortenedURL(req string) (*Link, error)
	GetAllKeys() ([]string, error)
	StoreKey(req *Link) error
	SetKeys(req []string) error
}

func (r *repository) Store(req *Link) error {
	m := new(memcache.Item)
	mapLinkToMemCacheItem(req, m)
	return r.db.Set(m)
}

func (r *repository) GetByShortenedURL(req string) (*Link, error) {
	l := new(Link)

	item, err := r.db.Get(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(item.Value, l)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (r *repository) StoreKey(req *Link) error {
	newKeys := new(keysCache)

	k, err := r.getKeys()
	if k == nil && err != nil {
		return err
	}

	if k != nil {
		newKeys = k
	}
	newKeys.Keys = append(newKeys.Keys, req.ShortURL)
	log.Println("XXX", newKeys.Keys)

	jsonData, _ := json.Marshal(newKeys)
	log.Println(string(jsonData))
	newItems := &memcache.Item{
		Key:   defaultKey,
		Value: jsonData,
	}

	return r.db.Set(newItems)
}

func (r *repository) GetAllKeys() ([]string, error) {
	k := new(keysCache)

	item, err := r.db.Get(defaultKey)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(item.Value, k)
	if err != nil {
		return nil, err
	}
	log.Println("test", k)

	return k.Keys, nil
}

func (r *repository) getKeys() (*keysCache, error) {
	keys := new(keysCache)

	item, err := r.db.Get(defaultKey)
	if err != nil && errors.Is(err, memcache.ErrCacheMiss) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	err = json.Unmarshal(item.Value, keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (r *repository) SetKeys(req []string) error {
	k := new(keysCache)
	k.Keys = req

	jsonData, _ := json.Marshal(k)
	newItems := &memcache.Item{
		Key:   "keys",
		Value: jsonData,
	}

	return r.db.Set(newItems)
}

func mapLinkToMemCacheItem(l *Link, m *memcache.Item) {
	jsonData, _ := json.Marshal(l)

	m.Key = l.ShortURL
	m.Value = jsonData
	if l.ExpireAt != nil {
		m.Expiration = l.GetTimeDiffInSecs()
	}
}
