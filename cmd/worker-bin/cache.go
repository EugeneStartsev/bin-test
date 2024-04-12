package main

import (
	"bin-checker/structs"
	"log"
	"sync"
)

type memCache struct {
	m     sync.RWMutex
	cache map[string]structs.BinData
}

func newMemCache() *memCache {
	return &memCache{
		cache: make(map[string]structs.BinData),
	}
}

func (c *memCache) set(key string, value structs.BinData) {
	c.m.Lock()
	c.cache[key] = value
	c.m.Unlock()
}

func (c *memCache) get(key string) (structs.BinData, bool) {
	c.m.RLock()
	val, ok := c.cache[key]
	c.m.RUnlock()

	if !ok {
		return structs.BinData{}, false
	}

	return val, true
}

func (c *memCache) recoverFromPostgres(stor storage) error {
	binData, err := stor.getAllBinsFromPostgres()
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range binData {
		c.set(val.Bin, val)
	}

	return nil
}
