package main

import (
	"bin-checker/structs"
	"log"
	"sync"
)

type memCache struct {
	m     sync.RWMutex
	cache map[string]structs.SaveBinData
}

func newMemCache() *memCache {
	return &memCache{
		cache: make(map[string]structs.SaveBinData),
	}
}

func (c *memCache) set(key string, value structs.SaveBinData) {
	c.m.Lock()
	c.cache[key] = value
	c.m.Unlock()
}

func (c *memCache) get(key string) (structs.SaveBinData, bool) {
	c.m.RLock()
	val, ok := c.cache[key]
	c.m.RUnlock()

	if !ok {
		return structs.SaveBinData{}, false
	}

	return val, true
}

func (c *memCache) recoverFromPostgres(stor storage) error {
	binData, err := stor.getAllBinsFromPostgres()
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range binData {
		c.set(val.Iin, val)
	}

	return nil
}
