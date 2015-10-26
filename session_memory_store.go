package xingyun

import (
	"time"

	"github.com/xiaoenai/gomemcache/memcache"
)

// Simple session storage using memory, handy for development
// **NEVER** use it in production!!!
type memoryStore struct {
	data map[string][]byte
}

func NewMemoryStore() *memoryStore {
	return &memoryStore{
		data: make(map[string][]byte),
	}
}

func (ms *memoryStore) SetSession(sessionID string, key string, data []byte) {
	ms.data[sessionID+key] = data
}

func (ms *memoryStore) GetSession(sessionID string, key string) []byte {
	data, _ := ms.data[sessionID+key]
	return data
}

func (ms *memoryStore) ClearSession(sessionID string, key string) {
	delete(ms.data, sessionID+key)
}

type memcacheStore struct {
	mc     *memcache.Client
	logger Logger
}

func NewMemcacheStore(addr string, logger Logger) *memcacheStore {
	return &memcacheStore{
		mc:     memcache.New(addr),
		logger: logger,
	}
}

func (ms *memcacheStore) SetSession(sessionID string, key string, data []byte) {
	_7days := 7 * 24 * int32(time.Hour.Seconds())
	err := ms.mc.Set(&memcache.Item{Key: sessionID + ":" + key, Value: data, Expiration: _7days})
	if err != nil {
		ms.logger.Errorf("SetSession %s", err)
	}
}

func (ms *memcacheStore) GetSession(sessionID string, key string) []byte {
	item, err := ms.mc.Get(sessionID + ":" + key)
	if err != nil {
		if err != memcache.ErrCacheMiss {
			ms.logger.Errorf("GetSession %s", err)
		}
		return nil
	}
	return item.Value
}

func (ms *memcacheStore) ClearSession(sessionID string, key string) {
	err := ms.mc.Delete(sessionID + ":" + key)
	if err != nil {
		ms.logger.Errorf("ClearSession %s", err)
	}
}
